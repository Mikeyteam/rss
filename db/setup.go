package db

import (
	"database/sql"
	"fmt"
)

var (
	createTable = `CREATE TABLE IF NOT EXISTS news (
		id int auto_increment primary key,
		title varchar(255),
		text TEXT,
		news_date varchar(255) 
	) DEFAULT CHARACTER SET UTF8;`
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func SetupDB() *sql.DB {
	var err error
	db, err := sql.Open("mysql", "root:root@/news")
	PanicIf(err)

	ctble, err := db.Query(createTable)
	PanicIf(err)
	fmt.Println("Table create successull", ctble)

	return db
}