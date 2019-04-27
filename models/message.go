package models

type Message struct {
	From     uint `json:"from"`
	To uint `json:"to"`
	Text string `json:"message"`
}