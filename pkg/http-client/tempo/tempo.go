package tempo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"tempo-redmine-sync/pkg/system"
)

const TEMPO_API_ENDPOINT = "https://api.tempo.io/core/3"

func GetMyWorklog() error {
	req, err := http.NewRequest("GET", TEMPO_API_ENDPOINT+"/worklogs/user/"+os.Getenv("JIRA_ACCOUNT_ID")+"?limit=1", nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("TEMPO_API_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Sending HTTP request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Parsing HTTP response: %v", err)
	}

	var myWorklog MyWorklog
	if err := json.Unmarshal(body, &myWorklog); err != nil {
		return fmt.Errorf("Parsing response to JSON: %v", err)
	}

	log.Println("[INFO] - Tempo autheticated user:", myWorklog.Results[0].Author.DisplayName)

	return nil
}

func PostWorklog(jobID int, act system.Activity) error {
	workLogBody := &WorkLogBody{
		IssueKey:         act.JiraIssue,
		TimeSpentSeconds: int(act.SpentTime * 3600),
		BillableSeconds:  int(act.SpentTime * 3600),
		StartDate:        act.Date,
		StartTime:        "08:00:00",
		Description:      act.Description,
		AuthorAccountID:  os.Getenv("JIRA_ACCOUNT_ID"),
	}

	reqBody, err := json.Marshal(workLogBody)
	if err != nil {
		return fmt.Errorf("Job[%d] Tempo Marshal body request: %v", jobID, err)
	}

	req, err := http.NewRequest("POST", TEMPO_API_ENDPOINT+"/worklogs", bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("TEMPO_API_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Job[%d] Tempo Sending HTTP request: %v", jobID, err)
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Job[%d] Tempo Parsing HTTP response: %v", jobID, err)
	}

	log.Printf("[INFO] - Job[%d] Tempo worklog created with %d status\n", jobID, resp.StatusCode)
	if resp.StatusCode != 200 {
		log.Printf("[INFO] - Job[%d] Redmine POST response: %s", jobID, resBody)
	}

	return nil
}
