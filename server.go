package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("./public/views/*.html")),
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.File("./public/views/index.html")
	})

	e.GET("/about-me", func(c echo.Context) error {
		return c.File("./public/views/about-me.html")
	})

	e.GET("/students", func(c echo.Context) error {
		return c.File("./public/views/students.html")
	})

	e.GET("/classes", func(c echo.Context) error {
		return c.File("./public/views/classes.html")
	})

	// TODO: add routes and lgoic for the following pages
	// student/{id}
	// class/{id}

	e.Logger.Fatal(e.Start(":1323"))
}
