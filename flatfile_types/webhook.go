package flatfile_types

type WebhookEvent struct {
	Type     string `json:"type"`
	ID       string `json:"id"`
	Sequence struct {
		Length int `json:"length"`
		Index  int `json:"index"`
	} `json:"sequence"`
}

type WebhookMeta struct {
	Batch struct {
		ID string `json:"id"`
	} `json:"batch"`
	Settings struct {
		ID *string `json:"id"`
	} `json:"settings"`
	Length int `json:"length"`
}

type WebhookData struct {
	Meta        WebhookMeta              `json:"meta"`
	ValidRows   []map[string]interface{} `json:"validRows"`
	InvalidRows []map[string]interface{} `json:"invalidRows"`
}

type WebhookBody struct {
	Event WebhookEvent `json:"event"`
	Data  WebhookData  `json:"data"`
}
