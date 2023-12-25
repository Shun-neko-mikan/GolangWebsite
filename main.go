package main

import (
	"net/http"

	"html/template"

	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}



func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 静的ファイルのルーティング
	e.Static("/static", "public/static")

	e.GET("/", viewMainPage)
	e.GET("/dtlPage", viewDetailPae)
	e.GET("/savegps", SaveGPS)
	e.GET("/getYoutube/:id", GetYoutube)
	e.Logger.Fatal(e.Start(":8080"))
}