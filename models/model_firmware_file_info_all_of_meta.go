package models

type FirmwareFileInfoAllOfMeta struct {

	// Firmware version in text format
	Version string `json:"version,omitempty"`

	PrinterType int32 `json:"printer_type,omitempty"`

	PrinterVersion int32 `json:"printer_version,omitempty"`
}
