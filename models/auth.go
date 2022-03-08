package models

type SignInp struct {
	Shortname string `json:"shortname"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
	Fullname  string `json:"fullname"`
}
