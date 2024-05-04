package components

import "github.com/taiki-nd/actions_notification/app/module/githubActions"

var GithubActionsClient *githubActions.GithubActionsClient

type githubActionsComponent struct{}

var GithubActions = &githubActionsComponent{}

func (githubActionsComponent) GetGithubActionsInfo() *githubActions.GithubActionsClient {
	return GithubActionsClient
}
