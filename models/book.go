package models

type Book struct {
	ID     int    `json:"ID"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Year   string `json:"Year"`
}
