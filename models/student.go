package models

type Student struct {
	LastName       string  `json:"last_name,omitempty"`
	FirstName      string  `json:"first_name,omitempty"`
	Section        Section `json:"section,omitempty"`
	SchoolDuration int     `json:"school_duration,omitempty"`
}
