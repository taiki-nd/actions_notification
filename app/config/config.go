package config

var GlobalConfig *Config

type Config struct {
	MessageApp     string
	DiscordWebhook string `json:"DISCORD_WEBHOOK" env:"DISCORD_WEBHOOK"`
	Repository     string `json:"GITHUB_REPOSITORY" env:"GITHUB_REPOSITORY"`
	SHA            string `json:"GITHUB_SHA" env:"GITHUB_SHA"`
	Ref            string `json:"GITHUB_REF" env:"GITHUB_REF"`
	Actor          string `json:"GITHUB_ACTOR" env:"GITHUB_ACTOR"`
	Workflow       string `json:"GITHUB_WORKFLOW" env:"GITHUB_WORKFLOW"`
	EventName      string `json:"GITHUB_EVENT_NAME" env:"GITHUB_EVENT_NAME"`
	Workspace      string `json:"GITHUB_WORKSPACE"`
	Branch         string `json:"GITHUB_BRANCH"`
	RunID          string `json:"GITHUB_RUN_ID"`
	ServerURL      string `json:"GITHUB_SERVER_URL"`
	Status         string `json:"GITHUB_ACTIONS_STATUS"`
	CommitMsg      string `json:"GITHUB_ACTIONS_COMMIT_MESSAGE"`
}

func NewConfig(messageApp string) *Config {
	return &Config{
		MessageApp: messageApp,
	}
}
