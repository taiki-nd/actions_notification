package app

import (
	"actionsnotification/app/config"
	"actionsnotification/app/module/discord"
	"actionsnotification/app/module/githubActions"
	"actionsnotification/app/module/slack"
	"actionsnotification/app/services/components"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp/v3"
)

var MessageApp string

type GitHubActionsParam struct {
	Webhook    string `json:"WEBHOOK_URL"`
	Repository string `json:"GITHUB_REPOSITORY"`
	SHA        string `json:"GITHUB_SHA"`
	Ref        string `json:"GITHUB_REF"`
	Actor      string `json:"GITHUB_ACTOR"`
	Workflow   string `json:"GITHUB_WORKFLOW"`
	EventName  string `json:"GITHUB_EVENT_NAME"`
	Workspace  string `json:"GITHUB_WORKSPACE"`
	Branch     string `json:"GITHUB_BRANCH"`
	RunID      string `json:"GITHUB_RUN_ID"`
	ServerURL  string `json:"GITHUB_SERVER_URL"`
	Status     string `json:"GITHUB_STATUS"`
	CommitMsg  string `json:"GITHUB_COMMIT_MESSAGE"`
	PrTitle    string `json:"GITHUB_PR_TITLE"`
	PrUrl      string `json:"GITHUB_PR_URL"`
}

func InitApp(env string) error {
	var err error
	// get params
	var webhook, repo, sha, ref, actor, workflow, eventName, workSpace, branch, runId, serverUrl, status, commitMsg, prTitle, prUrl string
	if env == "local" {
		webhook, repo, sha, ref, actor, workflow, eventName, workSpace, branch, runId, serverUrl, status, commitMsg, prTitle, prUrl, err = LoadParamsFromJson()
		if err != nil {
			return err
		}
	} else {
		webhook = os.Getenv("WEBHOOK_URL")
		if webhook == "" {
			fmt.Println("webhook URL is required. Please set webhook URL.")
			os.Exit(1)
		}
		repo = os.Getenv("GITHUB_REPOSITORY")
		sha = os.Getenv("GITHUB_SHA")
		ref = os.Getenv("GITHUB_REF")
		actor = os.Getenv("GITHUB_ACTOR")
		workflow = os.Getenv("GITHUB_WORKFLOW")
		eventName = os.Getenv("GITHUB_EVENT_NAME")
		workSpace = os.Getenv("GITHUB_WORKSPACE")
		branch = os.Getenv("GITHUB_REF")
		runId = os.Getenv("GITHUB_RUN_ID")
		serverUrl = os.Getenv("GITHUB_SERVER_URL")
		status = os.Getenv("GITHUB_STATUS")
		commitMsg = os.Getenv("GITHUB_COMMIT_MESSAGE")
		prTitle = os.Getenv("GITHUB_PR_TITLE")
		prUrl = os.Getenv("GITHUB_PR_URL")
	}

	messengerType := ""
	if strings.Contains(webhook, "discord") {
		messengerType = "discord"
	}
	if strings.Contains(webhook, "slack") {
		messengerType = "slack"
	}
	if messengerType == "" {
		return fmt.Errorf("webhook url must be discord or slack")
	}

	config.GlobalConfig = config.NewConfig(messengerType)
	if messengerType == "discord" {
		components.DiscordClient = discord.NewDiscord(webhook)
	}
	if messengerType == "slack" {
		components.SlackClient = slack.NewSlack(webhook)
	}
	components.GithubActionsClient = githubActions.NewGithubActions(repo, sha, ref, actor, workflow, eventName, workSpace, branch, runId, serverUrl, status, commitMsg, prTitle, prUrl)

	pp.Println(components.GithubActionsClient)

	return nil
}

func LoadParamsFromJson() (string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, error) {
	data, err := os.ReadFile("local.json")
	if err != nil {
		return "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", err
	}

	var env GitHubActionsParam
	err = json.Unmarshal(data, &env)
	if err != nil {
		return "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", err
	}

	webhook := env.Webhook
	repo := env.Repository
	sha := env.SHA
	ref := env.Ref
	actor := env.Actor
	workflow := env.Workflow
	eventName := env.EventName
	workSpace := env.Workspace
	branch := env.Branch
	runId := env.RunID
	serverUrl := env.ServerURL
	status := env.Status
	commitMsg := env.CommitMsg
	prTitle := env.PrTitle
	prUrl := env.PrUrl

	return webhook, repo, sha, ref, actor, workflow, eventName, workSpace, branch, runId, serverUrl, status, commitMsg, prTitle, prUrl, nil
}
