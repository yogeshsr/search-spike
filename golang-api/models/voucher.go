package models

type Voucher struct {
	Sponsor string `json:"sponsor"`
	Title   string `json:"title"`
	Desc    string `json:"description"`
}