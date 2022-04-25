package model

type Song struct {
	Length int    `json:"length,omitempty"`
	Artist string `json:"artist,omitempty"`
	Song   string `json:"song,omitempty"`
	Genre  string `json:"genre,omitempty"`
}
