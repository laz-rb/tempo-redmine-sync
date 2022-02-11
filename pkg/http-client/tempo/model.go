package tempo

type MyWorklog struct {
	Results []struct {
		Author struct {
			Self        string `json:"self"`
			AccountID   string `json:"accountId"`
			DisplayName string `json:"displayName"`
		} `json:"author"`
	} `json:"results"`
}

type WorkLogBody struct {
	IssueKey         string `json:"issueKey"`
	TimeSpentSeconds int    `json:"timeSpentSeconds"`
	BillableSeconds  int    `json:"billableSeconds"`
	StartDate        string `json:"startDate"`
	StartTime        string `json:"startTime"`
	Description      string `json:"description"`
	AuthorAccountID  string `json:"authorAccountId"`
}
