package pkg

type PaginationRequest struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}
