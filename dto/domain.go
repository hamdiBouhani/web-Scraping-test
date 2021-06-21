package dto

type Domain struct {
	MaxDepth int    `json:"max_depth"`
	URL      string `json:"url"`
}

/*
- HTML Version
- Page Title
- Headings count by level
- Amount of internal and external links
- Amount of inaccessible links
- If a page contains a login form
*/

type pageInfo struct {
	Links map[string]int `json:"links"`
}

type DomainResponce struct {
	HTMLVersion                    string           `json:"html_version"`
	PageTitle                      string           `json:"page_title"`
	Headings                       map[string]int64 `json:"headings"`
	ExternalAndInternalLinksAmount int64            `json:"external_internal_links_amount"`
	InaccessibleLinksAmount        int64            `json:"inaccessible_links_amount"`
	ContainsLoginForm              bool             `json:"contains_login_form"`
	PageInfo                       *pageInfo        `json:"page_info"`
}

func NewDomainResponce() *DomainResponce {
	return &DomainResponce{
		HTMLVersion: "",
		PageTitle:   "",
		Headings: map[string]int64{
			"h1": 0,
			"h2": 0,
			"h3": 0,
			"h4": 0,
			"h5": 0,
			"h6": 0,
		},
		ExternalAndInternalLinksAmount: 0,
		InaccessibleLinksAmount:        0,
		ContainsLoginForm:              false,
		PageInfo: &pageInfo{
			Links: make(map[string]int),
		},
	}
}
