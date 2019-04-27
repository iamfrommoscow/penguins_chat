package models

type Message struct {
	ID uint `json:"id"`
	From     uint `json:"from"`
	To uint `json:"to"`
	Text string `json:"message"`
}