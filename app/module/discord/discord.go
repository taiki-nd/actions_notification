package discord

type DiscordClient struct {
	WebhookUrl string
}

func NewDiscord(webhookUrl string) *DiscordClient {
	return &DiscordClient{
		WebhookUrl: webhookUrl,
	}
}
