package renderer

import (
	"fmt"
	"html/template"
	"io"
	"path"
)

type templateConfig struct {
	assetsPath    string
	viewsDir      string
	assetsMapping map[string]string
}

var HTML templateConfig = templateConfig{
	assetsPath:    "assets",
	viewsDir:      "internal/views",
	assetsMapping: map[string]string{},
}

func (tc *templateConfig) Configure(assetsPath string, viewsDir string, assetsMapping map[string]string) {
	tc.assetsPath = assetsPath
	tc.viewsDir = viewsDir
	tc.assetsMapping = assetsMapping
}

func (tc templateConfig) Render(writer io.Writer, data any, view string) {
	baseFile := path.Base(view)
	tmpl, err := template.New(baseFile).Funcs(template.FuncMap{
		"assetsPath": tc.getAssetsPath,
	}).ParseFiles(tc.viewsDir + "/" + view)

	if err != nil {
		fmt.Println("Error on rendering HTML:\n" + err.Error())
	}

	tmpl.Execute(writer, data)
}

func (tc templateConfig) getAssetsPath(filepath string) string {
	return HTML.assetsPath + "/" + HTML.assetsMapping[filepath]
}
