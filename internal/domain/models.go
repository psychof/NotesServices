package domain

type Notes struct {
	Title     string `validate:"required" json:"Title"`
	Text      string `validate:"required" json:"Text"`
	Time_zone string `validate:"required" json:"timeZone"`
}
