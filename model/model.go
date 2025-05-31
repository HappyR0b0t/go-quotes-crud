package model

type Quote struct {
	ID     uint   `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}
