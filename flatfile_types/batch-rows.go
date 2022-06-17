package flatfile_types

import "time"

type BatchRowsData struct {
	ID        string                 `json:"id"`
	BatchID   string                 `json:"batchId"`
	Raw       interface{}            `json:"raw"`
	Mapped    map[string]interface{} `json:"mapped"`
	Valid     bool                   `json:"valid"`
	Deleted   bool                   `json:"deleted"`
	Sequence  int                    `json:"sequence"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

type BatchRows struct {
	Pagination Pagination      `json:"pagination"`
	Data       []BatchRowsData `json:"data"`
}
