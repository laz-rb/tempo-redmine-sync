package main

import (
	"log"
	"os"
	"tempo-redmine-sync/pkg/http-client/redmine"
	"tempo-redmine-sync/pkg/http-client/tempo"
	"tempo-redmine-sync/pkg/system"

	"golang.org/x/sync/errgroup"
)

func main() {
	// Valid env vars are set
	if !system.EnvIsReady() {
		log.Println("[ERROR] - One ore more env variables are missing")
		os.Exit(1)
	}

	// Validate Redmine and Jira tokens
	log.Println("[INFO] - Validating user credentials...")
	eg := new(errgroup.Group)
	eg.Go(func() error {
		return redmine.GetMyAccount()
	})
	eg.Go(func() error {
		return tempo.GetMyWorklog()
	})

	if err := eg.Wait(); err != nil {
		log.Println("[ERROR] -", err)
		os.Exit(1)
	}

	//Iterate over user activities and post them to Tempo/Redmine
	userActivities, err := system.GetUserActivities()
	if err != nil {
		log.Println("[ERROR] - Couldn't get activities.json:", err)
		os.Exit(1)
	}

	for i, val := range userActivities.Activities {
		jobID := i
		activity := val
		eg.Go(func() error {
			return redmine.PostActivity(jobID, activity)
		})
		eg.Go(func() error {
			return tempo.PostWorklog(jobID, activity)
		})
	}

	if err := eg.Wait(); err != nil {
		log.Println("[ERROR] - Posting worklogs:", err)
		os.Exit(1)
	}
	log.Println("[INFO] - All jobs ended successfully!")
}
