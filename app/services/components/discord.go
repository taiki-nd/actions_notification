package components

import (
	"actionsnotification/app/module/discord"
	"actionsnotification/app/module/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/k0kubun/pp/v3"
)

var DiscordClient *discord.DiscordClient

type DiscordComponent struct {
	WebhookUrl string
	DiscordReq struct {
		Username  string   `json:"username"`
		AvatarURL string   `json:"avatar_url"`
		Content   string   `json:"content"`
		Embeds    []Embeds `json:"embeds"`
	}
}

type Embeds struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Timestamp   time.Time `json:"timestamp"`
	Color       string    `json:"color"`
	Footer      struct {
		Text    string `json:"text"`
		IconURL string `json:"icon_url"`
	} `json:"footer"`
	Image struct {
		URL string `json:"url"`
	} `json:"image"`
	Thumbnail struct {
		URL string `json:"url"`
	} `json:"thumbnail"`
	Author struct {
		Name    string `json:"name"`
		URL     string `json:"url"`
		IconURL string `json:"icon_url"`
	} `json:"author"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

func (discordComponent *DiscordComponent) MakeRequest() {
	actionsInfo := GithubActions.GetGithubActionsInfo()
	pp.Println(actionsInfo)

	var color string
	if actionsInfo.GithubActionsStatus.IsFailure() {
		color = "16711680"
	}
	if actionsInfo.GithubActionsStatus.IsSuccess() {
		// 緑
		color = "65280"
	}
	if actionsInfo.GithubActionsStatus.IsCancelled() {
		color = "16776960"
	}

	// set discord request
	discordComponent.WebhookUrl = DiscordClient.WebhookUrl
	discordComponent.DiscordReq.Username = "ActionsNotification"
	discordComponent.DiscordReq.AvatarURL = ""
	discordComponent.DiscordReq.Content = ""

	var embeds Embeds
	embeds.Title = fmt.Sprintf("%s: %s", actionsInfo.GithubActionsStatus.IconValue(), actionsInfo.GithubRepository)
	embeds.Description = fmt.Sprintf("by %s", actionsInfo.GithubActor)
	embeds.Timestamp = time.Now()
	embeds.Color = color
	discordComponent.DiscordReq.Embeds = append(discordComponent.DiscordReq.Embeds, embeds)

	var fields []Field
	workField := Field{
		Name:   "Workflow",
		Value:  fmt.Sprintf("[%s](<%s/%s/actions/run/%s>)", actionsInfo.GithubWorkflow, actionsInfo.GithubServerUrl, actionsInfo.GithubRepository, actionsInfo.GithubRunId),
		Inline: true,
	}
	repoField := Field{
		Name:   "Repository",
		Value:  fmt.Sprintf("[%s](<%s/%s>)", actionsInfo.GithubRepository, actionsInfo.GithubServerUrl, actionsInfo.GithubRepository),
		Inline: true,
	}
	braField := Field{
		Name:   "Branch",
		Value:  actionsInfo.GithubBranch,
		Inline: true,
	}
	var eveField Field
	if actionsInfo.GithubEventName.IsPullRequest() {
		eveField = Field{
			Name:   fmt.Sprintf("Event (%s)", actionsInfo.GithubEventName.UPPERValue()),
			Value:  fmt.Sprintf("PR URL: [%s](<%s>)", actionsInfo.GitHubActionsPrTitle, actionsInfo.GitHubActionsPrUrl),
			Inline: false,
		}
	} else {
		eveField = Field{
			Name:   fmt.Sprintf("Event (%s)", actionsInfo.GithubEventName),
			Value:  fmt.Sprintf("[%s](<%s/%s/commit/%s>): %s", utils.GetPrefix(actionsInfo.GithubSha, 7), actionsInfo.GithubServerUrl, actionsInfo.GithubRepository, actionsInfo.GithubSha, actionsInfo.GithubActionsCommitMsg),
			Inline: false,
		}
	}
	fields = append(fields, workField)
	fields = append(fields, repoField)
	fields = append(fields, braField)
	fields = append(fields, eveField)
	discordComponent.DiscordReq.Embeds[0].Fields = fields
}

func (discordComponent *DiscordComponent) SendRequest() error {
	// body
	json, err := json.Marshal(discordComponent.DiscordReq)
	if err != nil {
		fmt.Println(err)
		return err
	}

	body := bytes.NewBuffer(json)

	fmt.Println("discordComponent.WebhookUrl: ", discordComponent.WebhookUrl)
	fmt.Println("body: ", string(json))
	req, err := http.NewRequest("POST", discordComponent.WebhookUrl, body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 204 {
		fmt.Println("status code error: ", res.StatusCode)
	}

	return nil
}
