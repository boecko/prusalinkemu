package models

type Transfer struct {

	Type string `json:"type"`

	// Long Filename
	DisplayName string `json:"display_name"`

	Path string `json:"path"`

	Url string `json:"url,omitempty"`

	// Bytes
	Size string `json:"size,omitempty"`

	// Percents
	Progress float32 `json:"progress"`

	// Transferred data in bytes
	Transferred int32 `json:"transferred"`

	// Seconds
	TimeRemaining int32 `json:"time_remaining,omitempty"`

	// Seconds
	TimeTransferring int32 `json:"time_transferring"`

	// Whether or not print after finishing transfer (upload)
	ToPrint bool `json:"to_print"`
}
