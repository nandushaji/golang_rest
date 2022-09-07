package models

type Address struct {
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
}
