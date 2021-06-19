package rest

import (
	"web-Scraping-test/dto"
	"web-Scraping-test/pkg"

	"github.com/gin-gonic/gin"
)

func (RestServer) VisitUrls(ctx *gin.Context) {

	var domain dto.Domain
	if err := ctx.ShouldBind(&domain); err != nil {
		BindJsonErr(ctx, err)
		return
	}

	data := pkg.Crawl(domain.URL)

	ResponseData(ctx, data)

}

/*
	// Create a collector
	c := colly.NewCollector()

	// Set HTML callback
	// Won't be called if error occurs
	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	c.Visit(domain.URL)

*/
