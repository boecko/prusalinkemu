package models

// FirmwareFileInfoBasic - Simplified firmware file info within the folder's children
type FirmwareFileInfoBasic struct {

	// Short Filename
	Name string `json:"name"`

	ReadOnly bool `json:"read_only"`

	// Available for files only, not for folders
	Size int32 `json:"size,omitempty"`

	// File could be print file, firmware file, other (e.g. configuration) file, or folder
	Type string `json:"type"`

	// Timestamp in seconds
	MTimestamp int32 `json:"m_timestamp"`

	// Long Filename
	DisplayName string `json:"display_name,omitempty"`

	Refs FirmwareFileInfoBasicAllOfRefs `json:"refs,omitempty"`
}
