package models

type Error struct {

	// Prusa error code. Must be string if we will have printer with code for example 04
	Code string `json:"code,omitempty"`

	// Prusa error text string with prefilled variable macros.
	Title string `json:"title"`

	// Prusa error text string with prefilled variable macros.
	Text string `json:"text"`

	// Link to the Prusa help page
	Url string `json:"url,omitempty"`
}
