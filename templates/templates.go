package templates

import (
	"embed"
	"html/template"
	"io"
)

//go:embed default.tmpl
var DefaultTemplate string

//go:embed *
var files embed.FS

var (
	content = parse("content.tmpl")
	basic   = single("default.tmpl")
)

func Content(w io.Writer, data interface{}) error {
	return content.Execute(w, data)
}

func Basic(w io.Writer, data interface{}) error {
	return basic.Execute(w, data)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.tmpl").ParseFS(files, "layout.tmpl", file))
}

func single(file string) *template.Template {
	return template.Must(
		template.New(file).ParseFS(files, file))
}
