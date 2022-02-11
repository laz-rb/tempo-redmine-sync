package redmine

import (
	"time"
)

type MyAccount struct {
	User struct {
		ID          int       `json:"id"`
		Login       string    `json:"login"`
		Admin       bool      `json:"admin"`
		Firstname   string    `json:"firstname"`
		Lastname    string    `json:"lastname"`
		Mail        string    `json:"mail"`
		CreatedOn   time.Time `json:"created_on"`
		LastLoginOn time.Time `json:"last_login_on"`
	} `json:"user"`
}

type TimeEntryWrp struct {
	TimeEntry TimeEntry `json:"time_entry"`
}
type TimeEntry struct {
	IssueID    int     `json:"issue_id"`
	SpentOn    string  `json:"spent_on"`
	Hours      float64 `json:"hours"`
	ActivityID int     `json:"activity_id"`
	Comments   string  `json:"comments"`
}
