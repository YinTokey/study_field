package model

type Page struct {
	CurrentPage int    `json:"current_page"`
	TotalPages  int    `json:"total_pages"`
	TotalItems  int    `json:"total_items"`
	Feature     string `json:"feature"`
	Filters     struct {
	} `json:"filters"`
	Photos      []Photo `json:"photos"`

	
}
