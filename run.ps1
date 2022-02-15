docker run `
-e REDMINE_API_TOKEN=$env:REDMINE_API_TOKEN `
-e TEMPO_API_TOKEN=$env:TEMPO_API_TOKEN `
-e REDMINE_ISSUE_ID=$env:REDMINE_ISSUE_ID `
-e JIRA_ACCOUNT_ID=$env:JIRA_ACCOUNT_ID `
-v $env:ACTIVITIES_FILE:/app/activities.json `
tempo-redmine-sync