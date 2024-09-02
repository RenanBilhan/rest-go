package model

type Trip struct {
	tripId      int    `json:"Identification"`
	tripName    string `json:"Name"`
	description string `json:"Description"`
}
