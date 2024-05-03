package slack

type SlackClient struct {
	WebhookUrl string
}

func NewSlack(webhookUrl string) *SlackClient {
	return &SlackClient{
		WebhookUrl: webhookUrl,
	}
}
