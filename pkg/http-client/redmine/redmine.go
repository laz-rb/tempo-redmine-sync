package redmine

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

const REDMINE_ENDPOINT = "https://dev.unosquare.com/redmine"

func GetMyAccount() error {
	req, err := http.NewRequest("GET", REDMINE_ENDPOINT+"/my/account.json", nil)
	req.Header.Add("X-Redmine-API-Key", os.Getenv("REDMINE_API_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[ERROR] - Sending HTTP request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[ERROR] - Parsing HTTP response: %v", err)
	}

	var myAccount MyAccount
	if err := json.Unmarshal(body, &myAccount); err != nil {
		return fmt.Errorf("[ERROR] - Parsing to JSON: %v", err)
	}

	log.Println("[INFO] - Redmine authenticated user:", myAccount.User.Login)

	return nil
}

func PostActivity(jobID int, act system.Activity) error {
	redmineIssueID, _ := system.GetIntEnvVar()
	timeEntry := &TimeEntryWrp{
		TimeEntry: TimeEntry{
			IssueID:    redmineIssueID,
			SpentOn:    act.Date,
			Hours:      act.SpentTime,
			ActivityID: act.RedmineActivity,
			Comments:   act.Description,
		},
	}

	reqBody, err := json.Marshal(timeEntry)
	if err != nil {
		return fmt.Errorf("[ERROR] - Job[%d] Redmine Marshal body request: %v", jobID, err)
	}

	req, err := http.NewRequest("POST", REDMINE_ENDPOINT+"/time_entries.json", bytes.NewBuffer(reqBody))
	req.Header.Add("X-Redmine-API-Key", os.Getenv("REDMINE_API_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[ERROR] - Job[%d] Redmine Sending HTTP request: %v", jobID, err)
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[ERROR] - Job[%d] Redmine Parsing HTTP response: %v", jobID, err)
	}

	log.Printf("[INFO] - Job[%d] Redmine activity created with %d status\n", jobID, resp.StatusCode)
	if resp.StatusCode != 201 {
		log.Printf("[INFO] - Job[%d] Redmine POST response: %s", jobID, resBody)
	}

	return nil
}
