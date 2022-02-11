package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func EnvIsReady() bool {
	result := false
	envVars := [...]string{"REDMINE_API_TOKEN", "TEMPO_API_TOKEN", "REDMINE_ISSUE_ID", "JIRA_ACCOUNT_ID"}
	for _, val := range envVars {
		if _, res := os.LookupEnv(val); res {
			result = true
		} else {
			result = false
			break
		}
	}

	return result
}

func GetUserActivities() (*UserWork, error) {
	// Open our jsonFile
	jsonFile, err := os.Open("./activities.json")
	if err != nil {
		log.Println("[ERROR] - Opening JSON file:", err)
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("[ERROR] - Parsing JSON to byte:", err)
		return nil, err
	}

	userActivities := UserWork{}
	if err := json.Unmarshal(data, &userActivities); err != nil {
		return nil, fmt.Errorf("[ERROR] - Unmarshal to JSON: %v", err)
	}

	return &userActivities, nil
}

func GetIntEnvVar() (int, error) {
	res, err := strconv.Atoi(os.Getenv("REDMINE_ISSUE_ID"))
	if err != nil {
		log.Println("[ERROR] - Parsing env var to int:", err)
		return 0, err
	}
	return res, nil
}

type UserWork struct {
	Activities []Activity `json:"activities"`
}
type Activity struct {
	Date            string  `json:"date"`
	Description     string  `json:"description"`
	SpentTime       float64 `json:"spent_time"`
	JiraIssue       string  `json:"jira_issue"`
	RedmineActivity int     `json:"redmine_activity"`
}
