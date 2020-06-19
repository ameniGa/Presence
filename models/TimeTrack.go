package models

type TimeTrack struct {
	UserID string `json:"UserID"`
	Entry  string `json:"Entry"`
	Exit   string `json:"Exit"`
	Date   string `json:"Date"`
}
