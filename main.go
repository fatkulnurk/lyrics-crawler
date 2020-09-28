package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"log"
	"lyrics_crawler/abstract"
	"lyrics_crawler/sites"
	"strings"
)

func main() {
	urlVisits := []string{"https://www.chordtela.com/"}

	for i :=0; i < len(urlVisits); i++ {
		scape(urlVisits[i])
	}
}

func scape(urlVisit string)  {

	c := colly.NewCollector(
		colly.AllowedDomains("www.chordtela.com"),
	)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		var lyrics abtract.Lyric

		if strings.Contains(e.Request.URL.String(), "chordtela.com") {
			lyrics = sites.Chordtela(e)
		}

		if lyrics != (abtract.Lyric{}) {
			insertData(lyrics)
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Visit(urlVisit)
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lyrics_crawl")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func insertData(lyric abtract.Lyric)  {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into lyrics (title, body, url) values (?, ?, ?)", lyric.Title, lyric.Body, lyric.URL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Println("insert success!")
}