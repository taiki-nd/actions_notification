package components

import (
	"actionsnotification/app/module/discord"
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

	// set discord request
	discordComponent.WebhookUrl = DiscordClient.WebhookUrl
	discordComponent.DiscordReq.Username = "ActionsNotification"
	discordComponent.DiscordReq.AvatarURL = ""
	discordComponent.DiscordReq.Content = ""

	// actionsのsteps内にエラーがある場合は赤色にする
	color := "255"
	if actionsInfo.GithubActionsStatus == "failure" {
		color = "16711680"
	}

	var embeds Embeds
	embeds.Title = actionsInfo.GithubWorkflow
	embeds.Description = fmt.Sprintf("[workflow URL](<%s/%s/actions/runs/%s>)", actionsInfo.GithubServerUrl, actionsInfo.GithubRepository, actionsInfo.GithubRunId)
	embeds.Timestamp = time.Now()
	embeds.Color = "255"
	discordComponent.DiscordReq.Embeds = append(discordComponent.DiscordReq.Embeds, embeds)

	var fields []Field
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
	triField := Field{
		Name:   "Trigger",
		Value:  actionsInfo.GithubActor,
		Inline: true,
	}
	eveField := Field{
		Name:   "Event",
		Value:  actionsInfo.GithubEventName,
		Inline: false,
	}
	fields = append(fields, repoField)
	fields = append(fields, braField)
	fields = append(fields, triField)
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
