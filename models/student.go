package models

type Student struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Grade   string `json:"grade"`
}
