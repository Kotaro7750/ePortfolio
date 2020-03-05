package model

type Ticker struct {
	Id        int     `json:"id"`
	Ticker    string  `json:"ticker"`
	Devidened float64 `json:"devidened"`
}
