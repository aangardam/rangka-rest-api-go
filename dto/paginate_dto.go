package dto

type Paginate struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	Totalpage int `json:"total_page"`
}
