package flatfile_types

import "time"

type BatchHeader struct {
	Index  int    `json:"index"`
	Letter string `json:"letter"`
	Value  string `json:"value"`
}

type BatchHeaderMatched struct {
	Index      int    `json:"index"`
	Letter     string `json:"letter"`
	MatchedKey string `json:"matched_key"`
}

type BatchEndUser struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CompanyID   string    `json:"companyId"`
	CompanyName string    `json:"companyName"`
	TeamID      int       `json:"teamId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type BatchResponse struct {
	ID                   string                 `json:"id"`
	LicenseID            int                    `json:"licenseId"`
	TeamID               int                    `json:"teamId"`
	EnvironmentID        string                 `json:"environmentId"`
	WorkspaceID          *string                `json:"workspaceId"`
	SchemaID             *string                `json:"schemaId"`
	EmbedID              *string                `json:"embedId"`
	EndUserID            string                 `json:"endUserId"`
	Filename             string                 `json:"filename"`
	Source               *string                `json:"source"`
	Memo                 *string                `json:"memo"`
	ValidatedIn          *string                `json:"validatedIn"`
	OriginalFile         string                 `json:"originalFile"`
	Manual               bool                   `json:"manual"`
	Managed              bool                   `json:"managed"`
	CountRows            int                    `json:"countRows"`
	CountRowsInvalid     *int                   `json:"countRowsInvalid"`
	CountRowsAccepted    int                    `json:"countRowsAccepted"`
	CountColumns         int                    `json:"countColumns"`
	CountColumnsMatched  int                    `json:"countColumnsMatched"`
	HeadersRaw           []BatchHeader          `json:"headersRaw"`
	HeadersMatched       []BatchHeaderMatched   `json:"headersMatched"`
	Status               *string                `json:"status"`
	Type                 int                    `json:"type"`
	TargetSchema         *string                `json:"targetSchema"`
	ParsingConfig        map[string]interface{} `json:"parsingConfig"`
	ImportedFromURL      string                 `json:"importedFromUrl"`
	HeaderHash           string                 `json:"headerHash"`
	FailureReason        *string                `json:"failureReason"`
	MatchedAt            time.Time              `json:"matchedAt"`
	SubmittedAt          time.Time              `json:"submittedAt"`
	FailedAt             time.Time              `json:"failedAt"`
	HandledAt            time.Time              `json:"handledAt"`
	WriteAccessKey       string                 `json:"writeAccessKey"`
	LegacySettingsID     string                 `json:"legacySettingsId"`
	LegacyWebhookURL     *string                `json:"legacyWebhookUrl"`
	LegacyHasFieldHooks  bool                   `json:"legacyHasFieldHooks"`
	LegacyHasRecordHooks bool                   `json:"legacyHasRecordHooks"`
	CreatedAt            time.Time              `json:"createdAt"`
	UpdatedAt            time.Time              `json:"updatedAt"`
	DevMode              bool                   `json:"devMode"`
	Deleted              bool                   `json:"deleted"`
	Schema               *string                `json:"__schema__"`
	Embed                map[string]interface{} `json:"__embed__"`
	EndUser              BatchEndUser           `json:"__endUser__"`
	HasEndUser           bool                   `json:"__has_endUser__"`
}
