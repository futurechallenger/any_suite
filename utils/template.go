package utils

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template renders html templates
type Template struct {
	templates *template.Template
}

// Render renders html templates
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// return t.templates.ExecuteTemplate(w, name, data)
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

// NewTemplateEngine returns a `Template` instance
func NewTemplateEngine() echo.Renderer {
	return &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}
