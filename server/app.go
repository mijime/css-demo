package server

import (
	"path"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
)

type App struct {
	Engine *gin.Engine
	React  *React
}

func NewApp(opts ...AppOptions) *App {
	options := AppOptions{}
	for _, i := range opts {
		options = i
		break
	}
	options.init()

	if !options.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	engine.Use(func(c *gin.Context) {
		id, _ := uuid.NewV4()
		c.Set("uuid", id)
	})

	bfs := NewBinaryFileSystem("build/Release/assets")
	mt := multitemplate.New()
	tmpl, err := bfs.CreateTemplate("templates/index.tmpl")
	Must(err)
	mt.Add("templates/index.tmpl", tmpl)
	engine.HTMLRender = mt

	fs := bfs.CreateServer()

	react := NewReact(
		options.Debug,
		"build/Release/assets/app.bundle.js",
		"templates/index.tmpl",
		engine,
	)

	engine.Use(func(c *gin.Context) {
		ext := path.Ext(c.Request.URL.Path)

		if ext == "" {
			react.Handle(c)
			c.Abort()
			return
		}

		if bfs.Exists(c.Request.URL.Path) {
			fs.ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}

		return
	})

	return &App{
		Engine: engine,
		React:  react,
	}
}

func (a *App) Run(addr string) {
	a.Engine.Run(addr)
}

type AppOptions struct {
	Debug bool
}

func (ao *AppOptions) init() {}
