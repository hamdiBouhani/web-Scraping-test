package pkg

import (
	"time"
	"web-Scraping-test/dto"

	"github.com/gocolly/colly"
)

func Crawl(url string) *dto.DomainResponce {

	res := dto.NewDomainResponce()
	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{
		Delay:       1 * time.Second, // Set a delay between requests to these domains
		RandomDelay: 1 * time.Second, // Add an additional random delay
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		res.PageTitle = e.Text
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		res.Headings["h1"]++
	})

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		res.Headings["h2"]++
	})

	c.OnHTML("h3", func(e *colly.HTMLElement) {
		res.Headings["h3"]++
	})

	c.OnHTML("h4", func(e *colly.HTMLElement) {
		res.Headings["h4"]++
	})

	c.OnHTML("h5", func(e *colly.HTMLElement) {
		res.Headings["h5"]++
	})

	c.OnHTML("h6", func(e *colly.HTMLElement) {
		res.Headings["h6"]++
	})

	// count links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			res.PageInfo.Links[link]++
			res.ExternalAndInternalLinksAmount++
		}

	})

	c.Visit(url)

	return res
}
