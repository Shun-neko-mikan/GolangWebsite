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

type CommonData struct {
	Title string
}

func viewMainPage(c echo.Context) error {
	var common = CommonData{Title: "Main Page"}

	data := struct {
		CommonData
		Constent string
	}{
		CommonData: common,
		Constent:   "Hello, Test!",
	}
	return c.Render(http.StatusOK, "mainPage", data)
}

func viewDetailPae(c echo.Context) error {
	code := c.QueryParam("code")
	var common = CommonData{Title: "Detail Page"}

	data := struct {
		CommonData
		Code string
	}{
		CommonData: common,
		Code: code,
	}
	return c.Render(http.StatusOK, "detailPage", data)
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
	e.Logger.Fatal(e.Start(":8080"))
}