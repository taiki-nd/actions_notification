package components

import "actionsnotification/app/module/githubActions"

var GithubActionsClient *githubActions.GithubActionsClient

type githubActionsComponent struct{}

var GithubActions = &githubActionsComponent{}

func (githubActionsComponent) GetGithubActionsInfo() *githubActions.GithubActionsClient {
	return GithubActionsClient
}
