package pkg

import (
	"fmt"
	"strings"
	"time"
	"web-Scraping-test/dto"

	"github.com/PuerkitoBio/goquery"
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

	c.OnHTML("div", func(e *colly.HTMLElement) {
		e.DOM.Find("h1").Each(func(i int, s *goquery.Selection) {
			res.Headings["h1"]++
		})
		e.DOM.Find("h2").Each(func(i int, s *goquery.Selection) {
			res.Headings["h2"]++
		})
		e.DOM.Find("h3").Each(func(i int, s *goquery.Selection) {
			res.Headings["h3"]++
		})
		e.DOM.Find("h4").Each(func(i int, s *goquery.Selection) {
			res.Headings["h4"]++
		})
		e.DOM.Find("h5").Each(func(i int, s *goquery.Selection) {
			res.Headings["h5"]++
		})
		e.DOM.Find("h6").Each(func(i int, s *goquery.Selection) {
			res.Headings["h6"]++
		})
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Goquery selection of the HTMLElement is in e.DOM
		goquerySelection := e.DOM
		formTags := goquerySelection.Find("button")

		formTags.Each(func(_ int, s *goquery.Selection) {
			content := formTags.Text()
			nameProperty, _ := s.Attr("name")

			if strings.EqualFold(strings.ToLower(content), "sign in") ||
				strings.EqualFold(strings.ToLower(content), "signin") ||
				strings.EqualFold(strings.ToLower(content), "sign up") ||
				strings.EqualFold(strings.ToLower(content), "signup") ||
				strings.EqualFold(strings.ToLower(nameProperty), "sign in") ||
				strings.EqualFold(strings.ToLower(nameProperty), "signin") ||
				strings.EqualFold(strings.ToLower(nameProperty), "sign up") ||
				strings.EqualFold(strings.ToLower(nameProperty), "signup") {
				fmt.Printf("*** \n")
				res.ContainsLoginForm = true
			}

		})

	})

	c.Visit(url)

	return res
}
