package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func EnvIsReady() bool {
	result := false
	envVars := [...]string{"REDMINE_API_TOKEN", "TEMPO_API_TOKEN", "REDMINE_ISSUE_ID", "JIRA_ACCOUNT_ID", "POST_TO"}
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
	jsonFile, err := os.Open("/app/activities.json")
	if err != nil {
		return nil, fmt.Errorf("Opening JSON file: %v", err)
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("Parsing JSON to byte: %v", err)
	}

	userActivities := UserWork{}
	if err := json.Unmarshal(data, &userActivities); err != nil {
		return nil, fmt.Errorf("Unmarshal to JSON: %v", err)
	}

	return &userActivities, nil
}

func GetIntEnvVar(envVar string) (int, error) {
	res, err := strconv.Atoi(os.Getenv(envVar))
	if err != nil {
		return 0, fmt.Errorf("Parsing env var to int: %v", err)
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
