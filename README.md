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
 "url":"http://go-colly.org/"
}'
```

**Description**
**Input**

- URL (POST) : /apis/visit_url

```
{
 "url": "http://go-colly.org/"
}
```

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

# project structur

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
