package routes

import (
	"database/sql"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/mmcdole/gofeed"
	"net/http"
	"net/url"
	"rss/models"
	s "strings"
)

func HomeRouterHandler(render render.Render) {
	render.HTML(200, "index", nil)
}

func ViewRouterHandler(render render.Render, db *sql.DB) {
	rows, err := db.Query("select * from news")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	news := []models.News{}

	for rows.Next() {
		newsModel := models.News{}
		err := rows.Scan(&newsModel.Id, &newsModel.Title, &newsModel.NewsDate, &newsModel.Text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		news = append(news, newsModel)
	}

	render.HTML(200, "view", news)
}

func CreateRouteHandler(render render.Render) {
	render.HTML(200, "create", nil)
}

func DeleteRouteHandler(render render.Render, w http.ResponseWriter, r *http.Request, params martini.Params, db *sql.DB) {
	id := params["id"]
	if id == "all" {
		db.Query("DELETE from news")
	} else {
		rows, err := db.Query("DELETE from news Where id = " + id)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
	}

	render.Redirect("/view")
}

func SafeRssfHandler(render render.Render, r *http.Request, db *sql.DB) {
	linkRss := r.FormValue("linkRss")
	ruleParse := r.FormValue("ruleParse")

	_, errLong := url.ParseRequestURI(linkRss)
	if errLong != nil {
		render.Redirect("/create")
		return
	}

	parser := gofeed.NewParser()
	news, err := parser.ParseURL(linkRss)

	if err != nil {
		fmt.Println("Нет новостей")

		render.Redirect("/create")
	} else {
		for _, row := range news.Items {
			if s.Contains(row.Title, ruleParse) {
				db.Exec("insert into news (title, text, news_date) values (?, ?, ?)",
					row.Title, row.Description, row.Published)
			} else if ruleParse == "" {
				db.Exec("insert into news (title, text, news_date) values (?, ?, ?)",
					row.Title, row.Description, row.Published)
			}
		}

		render.Redirect("/view")
	}
}

func SearchHandler(render render.Render, r *http.Request, db *sql.DB) {
	searchRss := r.FormValue("search")
	rows, err := db.Query("SELECT * FROM `news` WHERE `title` LIKE concat('%', ?, '%')", searchRss)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	news := []models.News{}

	for rows.Next() {
		newsModel := models.News{}
		err := rows.Scan(&newsModel.Id, &newsModel.Title, &newsModel.NewsDate, &newsModel.Text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		news = append(news, newsModel)
	}

	render.HTML(200, "view", news)

}
