package model

type Trip struct {
	TripId      int    `json:"Identification"`
	TripName    string `json:"Name"`
	Description string `json:"Description"`
}
