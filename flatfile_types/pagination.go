package flatfile_types

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	Limit       int `json:"limit"`
	Offset      int `json:"offset"`
	OnPage      int `json:"onPage"`
	TotalCount  int `json:"totalCount"`
	PageCount   int `json:"pageCount"`
}
