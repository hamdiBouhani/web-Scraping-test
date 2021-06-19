package rest

import (
	"web-Scraping-test/dto"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func (RestServer) VisitUrls(ctx *gin.Context) {

	var domains dto.Domains
	if err := ctx.ShouldBind(&domains); err != nil {
		BindJsonErr(ctx, err)
		return
	}

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	fmt.Println(c)

	ResponseData(ctx, domains)

}
