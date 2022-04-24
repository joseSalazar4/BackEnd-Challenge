package model

type Song struct {
	Length int    `json:"length"`
	Artist string `json:"artist"`
	Song   string `json:"song"`
	Genre  string `json:"genre"`
}
