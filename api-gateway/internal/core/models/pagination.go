package models

type PaginationLinks struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}


type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	LastPage    int `json:"last_page"`
	Path        string `json:"path"`
	PerPage     int `json:"per_page"`
	To          int `json:"to"`
	Total       int `json:"total"`
}
