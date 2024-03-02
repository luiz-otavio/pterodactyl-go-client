package http

type PteroRequestSchema struct {
	Object     string   `json:"object"`
	Attributes struct{} `json:"attributes"`
}

type PteroMetaSchema struct {
	Meta struct{} `json:"meta"`
}

type PteroPaginationSchema struct {
	Pagination struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
	}
}
