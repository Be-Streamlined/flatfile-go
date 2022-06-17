package flatfile_types

import "time"

type ListBatches_Batch struct {
	ID                   string        `json:"id"`
	LicenseID            int           `json:"licenseId"`
	TeamID               int           `json:"teamId"`
	EnvironmentID        string        `json:"environmentId"`
	WorkspaceID          *string       `json:"workspaceId"`
	SchemaID             *string       `json:"schemaId"`
	EmbedID              *string       `json:"embedId"`
	EndUserID            string        `json:"endUserId"`
	Filename             string        `json:"filename"`
	Source               *string       `json:"source"`
	Memo                 *string       `json:"memo"`
	ValidatedIn          *string       `json:"validatedIn"`
	OriginalFile         string        `json:"originalFile"`
	Manual               bool          `json:"manual"`
	Managed              bool          `json:"managed"`
	CountRows            int           `json:"countRows"`
	CountRowsInvalid     *int          `json:"countRowsInvalid"`
	CountRowsAccepted    int           `json:"countRowsAccepted"`
	CountColumns         int           `json:"countColumns"`
	CountColumnsMatched  int           `json:"countColumnsMatched"`
	HeadersRaw           []BatchHeader `json:"headersRaw"`
	HeadersMatched       []BatchHeader `json:"headersMatched"`
	Status               *string       `json:"status"`
	Type                 int           `json:"type"`
	TargetSchema         *string       `json:"targetSchema"`
	ParsingConfig        *string       `json:"parsingConfig"`
	ImportedFromURL      string        `json:"importedFromUrl"`
	HeaderHash           string        `json:"headerHash"`
	FailureReason        *string       `json:"failureReason"`
	MatchedAt            time.Time     `json:"matchedAt"`
	SubmittedAt          time.Time     `json:"submittedAt"`
	FailedAt             time.Time     `json:"failedAt"`
	HandledAt            time.Time     `json:"handledAt"`
	WriteAccessKey       string        `json:"writeAccessKey"`
	LegacySettingsID     string        `json:"legacySettingsId"`
	LegacyWebhookURL     *string       `json:"legacyWebhookUrl"`
	LegacyHasFieldHooks  bool          `json:"legacyHasFieldHooks"`
	LegacyHasRecordHooks bool          `json:"legacyHasRecordHooks"`
	CreatedAt            time.Time     `json:"createdAt"`
	UpdatedAt            time.Time     `json:"updatedAt"`
	DevMode              bool          `json:"devMode"`
	Deleted              bool          `json:"deleted"`
}

type ListBatchesResponse struct {
	Pagination Pagination          `json:"pagination"`
	Data       []ListBatches_Batch `json:"data"`
}
