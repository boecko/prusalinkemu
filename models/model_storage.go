package models

type Storage struct {

	// Name of the storage, based on selected language
	Name string `json:"name,omitempty"`

	// Storage source
	Type string `json:"type"`

	// Path to storage (not display path)
	Path string `json:"path"`

	// Size of all print files in bytes
	PrintFiles int32 `json:"print_files,omitempty"`

	// Size of all system files in bytes
	SystemFiles int32 `json:"system_files,omitempty"`

	// System free space in bytes
	FreeSpace int32 `json:"free_space,omitempty"`

	// System total space in bytes
	TotalSpace int32 `json:"total_space,omitempty"`

	// Whether the storage is available or not
	Available bool `json:"available"`

	// Whether the storage is read only
	ReadOnly bool `json:"read_only,omitempty"`
}
