package main

import (
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/render"
	"rss/db"
	"rss/routes"
)

func main() {
	m := martini.Classic()
	m.Map(db.SetupDB())

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	staticOption := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOption))

	m.Get("/", routes.HomeRouterHandler)
	m.Get("/view", routes.ViewRouterHandler)
	m.Get("/create", routes.CreateRouteHandler)
	m.Post("/saveRss", routes.SafeRssfHandler)
	m.Post("/search", routes.SearchHandler)
	m.Get("/delete/:id", routes.DeleteRouteHandler)

	m.Run()

}
