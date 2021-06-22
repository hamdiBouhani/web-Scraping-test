# web-Scraping-test

Create a web application which takes a website URL as an input and provides general information
about the contents of the page:

- HTML Version
- Page Title
- Headings count by level
- Amount of internal and external links
- Amount of inaccessible links
- If a page contains a login form

To start project

```
make -B start-app

or

go run cmd/web-Scraping-test/main.go serve

```

# Apis

## Visit Url

```
curl --request POST \
  --url http://localhost:8080/apis/visit_url \
  --header 'content-type: application/json' \
  --data '{
 "max_depth": 2,
 "url": "https://www.google.com/"
}'
```

**Description**
**Input**

- URL (POST) : /apis/visit_url

```
{
 "max_depth": 1,
 "url": "http://go-colly.org/"
}
```

 // MaxDepth limits the recursion depth of visited URLs.

 // Set it to 0 for infinite recursion (default).

 //url ==> visited url

**Output**

- http status : 200 ok

```
{
  "data": {
    "html_version": "",
    "page_title": "Scraping Framework for Golang",
    "headings": {
      "h1": 0,
      "h2": 6,
      "h3": 0,
      "h4": 1,
      "h5": 0,
      "h6": 0
    },
    "external_internal_links_amount": 25,
    "inaccessible_links_amount": 0,
    "contains_login_form": false,
    "page_info": {
      "links": {
        "http://go-colly.org/": 3,
        "http://go-colly.org/articles/": 2,
        "http://go-colly.org/contact/": 1,
        "http://go-colly.org/datasets/": 2,
        "http://go-colly.org/docs/": 4,
        "http://go-colly.org/services/": 3,
        "http://go-colly.org/sitemap.xml": 1,
        "https://github.com/gocolly/colly": 5,
        "https://github.com/gocolly/colly/blob/master/LICENSE.txt": 1,
        "https://github.com/gocolly/site/": 1,
        "https://godoc.org/github.com/gocolly/colly": 2
      }
    }
  },
  "success": true
}
```

# Project Structure

```
.
├── cmd
│   └── web-Scraping-test
│       └── main.go
├── dto
│   └── domain.go
├── _exemple
│   ├── colly_basic_exemple.go
│   └── link_exemple.go
├── go.mod
├── go.sum
├── Makefile
├── pkg
│   └── crawl.go
├── README.md
└── svc
    ├── cmd
    │   └── serve
    │       └── serve.go
    ├── configs
    │   └── configs.go
    └── rest
        ├── handlers.go
        ├── response.go
        └── server.go

10 directories, 14 files
```

```
package pkg

import (
 "time"
 "web-Scraping-test/dto"

 "github.com/PuerkitoBio/goquery"
 "github.com/gocolly/colly"
 "github.com/sirupsen/logrus"
)

func Crawl(Log *logrus.Logger, domain dto.Domain) *dto.DomainResponce {

 res := dto.NewDomainResponce()
 c := colly.NewCollector()

 // MaxDepth limits the recursion depth of visited URLs.
 // Set it to 0 for infinite recursion (default).
 c.MaxDepth = domain.MaxDepth
 c.Limit(&colly.LimitRule{
  Delay:       1 * time.Second, // Set a delay between requests to these domains
  RandomDelay: 1 * time.Second, // Add an additional random delay
 })

 c.OnHTML("html", func(e *colly.HTMLElement) {
  xmlnsProperty := e.Attr("xmlns")
  if len(xmlnsProperty) > 0 {
   res.HTMLVersion = "< 5"
  }
  res.HTMLVersion = "5"
 })

 c.OnHTML("title", func(e *colly.HTMLElement) {
  res.PageTitle = e.Text
 })

 //input[type=password]
 c.OnHTML("input[type] ", func(e *colly.HTMLElement) {
  typeProperty := e.Attr("type")
  if typeProperty == "password" {
   res.ContainsLoginForm = true
  }
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

 // Find and visit all links
 c.OnHTML("a[href]", func(e *colly.HTMLElement) {
  e.Request.Visit(e.Attr("href"))
 })

 c.OnRequest(func(r *colly.Request) {
  Log.Println("Visiting", r.URL)
 })

 c.OnError(func(_ *colly.Response, err error) {
  Log.Errorln("Something went wrong:", err)
  res.InaccessibleLinksAmount++
 })

 c.Visit(domain.URL)

 return res
}

```
