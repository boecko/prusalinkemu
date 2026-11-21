package models

type Version struct {

	Api string `json:"api"`

	Version string `json:"version"`

	Printer string `json:"printer"`

	Text string `json:"text"`

	Firmware string `json:"firmware"`

	Sdk string `json:"sdk,omitempty"`

	Capabilities VersionCapabilities `json:"capabilities,omitempty"`
}
