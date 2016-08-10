package server

import (
	"html/template"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

type BinaryFileSystem struct {
	Root string
	Fs   http.FileSystem
}

func NewBinaryFileSystem(root string) *BinaryFileSystem {
	fs := &assetfs.AssetFS{Asset, AssetDir, AssetInfo, root}
	return &BinaryFileSystem{Fs: fs, Root: root}
}

func (b *BinaryFileSystem) Open(name string) (http.File, error) {
	return b.Fs.Open(name)
}

func (b *BinaryFileSystem) Exists(filepath string) bool {
	_, err := Asset(b.Root + filepath)
	return err == nil
}

func (b *BinaryFileSystem) CreateServer() http.Handler {
	return http.FileServer(b.Fs)
}

func (b *BinaryFileSystem) CreateTemplate(path string) (*template.Template, error) {
	tmplStr, err := Asset(path)

	if err != nil {
		return nil, err
	}

	return template.New(path).Parse(string(tmplStr))
}
