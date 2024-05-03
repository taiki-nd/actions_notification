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
)

var MessageApp string

func InitApp() error {
	// get params
	webhookUrl := os.Getenv("WEBHOOK_URL")
	if webhookUrl == "" {
		fmt.Println("webhook URL is required. Please set webhook URL.")
		os.Exit(1)
	}
	repo := os.Getenv("GITHUB_REPOSITORY")
	sha := os.Getenv("GITHUB_SHA")
	ref := os.Getenv("GITHUB_REF")
	actor := os.Getenv("GITHUB_ACTOR")
	workflow := os.Getenv("GITHUB_WORKFLOW")
	eventName := os.Getenv("GITHUB_EVENT_NAME")
	workSpace := os.Getenv("GITHUB_WORKSPACE")
	branch := os.Getenv("GITHUB_REF")
	runId := os.Getenv("GITHUB_RUN_ID")
	severUrl := os.Getenv("GITHUB_SERVER_URL")
	commitMsg := os.Getenv("GITHUB_ACTIONS_COMMIT_MESSAGE")

	// github actions statusを取得
	status := os.Getenv("GITHUB_ACTIONS_STATUS")

	messengerType := ""
	if strings.Contains(webhookUrl, "discord") {
		messengerType = "discord"
	}
	if strings.Contains(webhookUrl, "slack") {
		messengerType = "slack"
	}
	if messengerType == "" {
		return fmt.Errorf("webhook url must be discord or slack")
	}

	config.GlobalConfig = config.NewConfig(messengerType)
	if messengerType == "discord" {
		components.DiscordClient = discord.NewDiscord(webhookUrl)
	}
	if messengerType == "slack" {
		components.SlackClient = slack.NewSlack(webhookUrl)
	}
	components.GithubActionsClient = githubActions.NewGithubActions(repo, sha, ref, actor, workflow, eventName, workSpace, branch, runId, severUrl, status, commitMsg)

	return nil
}

type Env struct {
	DiscordWebhook string `json:"DISCORD_WEBHOOK"`
	Repository     string `json:"GITHUB_REPOSITORY"`
	SHA            string `json:"GITHUB_SHA"`
	Ref            string `json:"GITHUB_REF"`
	Actor          string `json:"GITHUB_ACTOR"`
	Workflow       string `json:"GITHUB_WORKFLOW"`
	EventName      string `json:"GITHUB_EVENT_NAME"`
	Workspace      string `json:"GITHUB_WORKSPACE"`
	Branch         string `json:"GITHUB_BRANCH"`
	RunID          string `json:"GITHUB_RUN_ID"`
	ServerURL      string `json:"GITHUB_SERVER_URL"`
	Status         string `json:"GITHUB_ACTIONS_STATUS"`
	CommitMsg      string `json:"GITHUB_ACTIONS_COMMIT_MESSAGE"`
}

func InitAppOnLocal() error {
	data, err := os.ReadFile("local.json")
	if err != nil {
		return err
	}

	var env Env
	err = json.Unmarshal(data, &env)
	if err != nil {
		return err
	}

	discordWebhook := env.DiscordWebhook
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

	config.GlobalConfig = config.NewConfig("discord")
	components.DiscordClient = discord.NewDiscord(discordWebhook)
	components.GithubActionsClient = githubActions.NewGithubActions(repo, sha, ref, actor, workflow, eventName, workSpace, branch, runId, serverUrl, status, commitMsg)

	return nil
}
