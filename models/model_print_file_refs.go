package models

// PrintFileRefs - Reference links for file thumbnail, icon and download
type PrintFileRefs struct {

	Download string `json:"download,omitempty"`

	Icon string `json:"icon,omitempty"`

	Thumbnail string `json:"thumbnail,omitempty"`
}
