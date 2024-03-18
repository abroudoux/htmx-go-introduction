package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("../views/*.html")),
	}
}

type Contact struct {
	Name string
	Email string
}

func newContact(name string, email string) Contact {
	return Contact{
		Name: name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: Contacts{
			newContact("John Doe", "jdhn.doe@gmail.com"),
			newContact("Jane Doe", "jane.doe@gmail.com"),
		},
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	data := newData()
	e.Renderer = NewTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		data.Contacts = append(data.Contacts, newContact(name, email))

		return c.Render(200, "display-contacts", data)
	})

	e.Logger.Fatal(e.Start(":42069"))
}