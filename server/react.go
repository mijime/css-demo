package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	"gopkg.in/olebedev/go-duktape-fetch.v2"
	"gopkg.in/olebedev/go-duktape.v2"
)

type React struct {
	pool
	debug    bool
	path     string
	template string
}

func NewReact(debug bool, path string, template string, engine http.Handler) *React {
	r := &React{
		debug:    debug,
		path:     path,
		template: template,
	}
	if debug {
		r.pool = newOnDemandPool(path, engine)
	} else {
		r.pool = newDuktapePool(path, engine, runtime.NumCPU()+1)
	}

	return r
}

func (r *React) Handle(c *gin.Context) {
	uuidContainer, _ := c.Get("uuid")
	UUID := uuidContainer.(*uuid.UUID)

	defer func() {
		if err := recover(); err != nil {
			c.HTML(http.StatusInternalServerError, r.template, gin.H{
				"UUID":  UUID.String(),
				"Error": err,
			})
		}
	}()

	vm := r.get()

	start := time.Now()

	select {
	case re := <-vm.Handle(map[string]interface{}{
		"url":     c.Request.URL.String(),
		"headers": c.Request.Header,
		"uuid":    UUID.String(),
	}):
		re.RenderTime = time.Since(start)

		r.put(vm)

		if len(re.Redirect) == 0 && len(re.Error) == 0 {
			c.Header("X-React-Render-Time", fmt.Sprintf("%s", re.RenderTime))
			c.HTML(http.StatusOK, r.template, re)

		} else if len(re.Redirect) != 0 {
			c.Redirect(http.StatusMovedPermanently, re.Redirect)

		} else if len(re.Error) != 0 {
			c.Header("X-React-Render-Time", fmt.Sprintf("%s", re.RenderTime))
			c.HTML(http.StatusInternalServerError, r.template, re)
		}
	case <-time.After(2 * time.Second):
		r.drop(vm)

		c.HTML(http.StatusInternalServerError, r.template, Resp{
			UUID:  UUID.String(),
			Error: "time is out",
		})
	}
}

type Resp struct {
	UUID       string        `json:"uuid"`
	Error      string        `json:"error"`
	Redirect   string        `json:"redirect"`
	App        string        `json:"app"`
	Title      string        `json:"title"`
	Meta       string        `json:"meta"`
	Initial    string        `json:"initial"`
	RenderTime time.Duration `json:"-"`
}

func (r Resp) HTMLApp() template.HTML {
	return template.HTML(r.App)
}

func (r Resp) HTMLTitle() template.HTML {
	return template.HTML(r.Title)
}

func (r Resp) HTMLMeta() template.HTML {
	return template.HTML(r.Meta)
}

type pool interface {
	get() *ReactVM
	put(*ReactVM)
	drop(*ReactVM)
}

func newReactVM(path string, engine http.Handler) *ReactVM {

	vm := &ReactVM{
		Context: duktape.New(),
		ch:      make(chan Resp, 1),
	}

	vm.PevalString(`var console = {log:print,warn:print,error:print,info:print}`)
	fetch.PushGlobal(vm.Context, engine)
	app, err := Asset(path)
	Must(err)

	vm.PushGlobalGoFunction("__goServerCallback__", func(ctx *duktape.Context) int {
		result := ctx.SafeToString(-1)
		vm.ch <- func() Resp {
			var re Resp
			json.Unmarshal([]byte(result), &re)
			return re
		}()
		return 0
	})

	fmt.Printf("%s loaded\n", path)
	if err := vm.PevalString(string(app)); err != nil {
		derr := err.(*duktape.Error)
		panic(derr.Message)
	}
	vm.PopN(vm.GetTop())
	return vm
}

type ReactVM struct {
	*duktape.Context
	ch chan Resp
}

func (r *ReactVM) Handle(req map[string]interface{}) <-chan Resp {
	b, err := json.Marshal(req)
	Must(err)

	r.PevalString(`main(` + string(b) + `, __goServerCallback__)`)
	return r.ch
}

func (r *ReactVM) DestroyHeap() {
	close(r.ch)
	r.Context.DestroyHeap()
}

func newOnDemandPool(path string, engine http.Handler) *onDemandPool {
	return &onDemandPool{
		path:   path,
		engine: engine,
	}
}

type onDemandPool struct {
	path   string
	engine http.Handler
}

func (f *onDemandPool) get() *ReactVM {
	return newReactVM(f.path, f.engine)
}

func (f onDemandPool) put(c *ReactVM) {
	c.Lock()
	c.FlushTimers()
	c.Gc(0)
	c.DestroyHeap()
}

func (f *onDemandPool) drop(c *ReactVM) {
	f.put(c)
}

func newDuktapePool(path string, engine http.Handler, size int) *duktapePool {
	pool := &duktapePool{
		path:   path,
		ch:     make(chan *ReactVM, size),
		engine: engine,
	}

	go func() {
		for i := 0; i < size; i++ {
			pool.ch <- newReactVM(path, engine)
		}
	}()

	return pool
}

type duktapePool struct {
	ch     chan *ReactVM
	path   string
	engine http.Handler
}

func (o *duktapePool) get() *ReactVM {
	return <-o.ch
}

func (o *duktapePool) put(ot *ReactVM) {
	ot.Lock()
	ot.FlushTimers()
	ot.Unlock()
	o.ch <- ot
}

func (o *duktapePool) drop(ot *ReactVM) {
	ot.Lock()
	ot.FlushTimers()
	ot.Gc(0)
	ot.DestroyHeap()
	ot = nil
	o.ch <- newReactVM(o.path, o.engine)
}
