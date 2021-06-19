package dto

type Domain struct {
	URL string `json:"url"`
}

/*
- HTML Version
- Page Title
- Headings count by level
- Amount of internal and external links
- Amount of inaccessible links
- If a page contains a login form
*/

// type Heading struct {
// 	Level string `json:"level"`
// 	Count int64  `json:"count"`
// }

type DomainResponce struct {
	HTMLVersion             string           `json:"html_version"`
	PageTitle               string           `json:"page_title"`
	Headings                map[string]int64 `json:"headings"`
	ExternalLinksAmount     int64            `json:"external_links_amount"`
	InternalLinksAmount     int64            `json:"internal_links_amount"`
	InaccessibleLinksAmount int64            `json:"inaccessible_links_amount"`
	ContainsLoginForm       bool             `json:"contains_login_form"`
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
		ExternalLinksAmount:     0,
		InternalLinksAmount:     0,
		InaccessibleLinksAmount: 0,
		ContainsLoginForm:       false,
	}
}
