package sites

import (
	"github.com/gocolly/colly"
	"lyrics_crawler/abtract"
	"strings"
)

func Chordtela(e *colly.HTMLElement) abtract.Lyric {
	var lyric abtract.Lyric
	title := e.ChildText("h1")
	title = strings.ReplaceAll(title, "Chord Kunci Gitar ", "")
	body := e.ChildText("pre")

	if body == "" {
		body = e.ChildText("div[class=post-body]")
	}

	url := e.Request.URL.String()

	if title != "" && body != "" {
		lyric.Title = title
		lyric.Body = body
		lyric.URL = url

		return lyric
	}

	return lyric
}