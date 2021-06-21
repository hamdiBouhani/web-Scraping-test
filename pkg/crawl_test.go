package pkg

import (
	"strings"
	"testing"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
)

// Borrowed from http://infohost.nmt.edu/tcc/help/pubs/xhtml/example.html
// Added attributes to the `<li>` tags for testing purposes
const htmlPage = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"
 "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
  <head>
    <title>Your page title here</title>
  </head>
  <body>
    <h1>Your major heading here</h1>
    <p>
      This is a regular text paragraph.
    </p>
    <ul>
      <li class="list-item-1">
        First bullet of a bullet list.
      </li>
      <li class="list-item-2">
        This is the <em>second</em> bullet.
      </li>
    </ul>
  </body>
</html>
`

func TestAttr(t *testing.T) {
	resp := &colly.Response{StatusCode: 200, Body: []byte(htmlPage)}
	doc, _ := htmlquery.Parse(strings.NewReader(htmlPage))
	xmlNode := htmlquery.FindOne(doc, "/html")
	xmlElem := colly.NewXMLElementFromHTMLNode(resp, xmlNode)

	if xmlElem.Attr("xmlns") != "http://www.w3.org/1999/xhtml" {
		t.Fatalf("failed xmlns attribute test: %v != http://www.w3.org/1999/xhtml", xmlElem.Attr("xmlns"))
	}

	if xmlElem.Attr("xml:lang") != "en" {
		t.Fatalf("failed lang attribute test: %v != en", xmlElem.Attr("lang"))
	}
}
