package githubActions

import "actionsnotification/app/core/consts"

type GithubActionsClient struct {
	GithubRepository       string
	GithubSha              string
	GithubRef              string
	GithubActor            string
	GithubWorkflow         string
	GithubEventName        consts.EventName
	GithubWorkSpace        string
	GithubBranch           string
	GithubRunId            string
	GithubServerUrl        string
	GithubActionsStatus    consts.JobStatus
	GithubActionsCommitMsg string
	GitHubActionsPrTitle   string
	GitHubActionsPrUrl     string
}

func NewGithubActions(repo, sha, ref, actor, workflow, eventName, workSpace, branch, runId, severUrl, status, commitMsg, prTitle, prUrl string) *GithubActionsClient {
	return &GithubActionsClient{
		GithubRepository:       repo,
		GithubSha:              sha,
		GithubRef:              ref,
		GithubActor:            actor,
		GithubWorkflow:         workflow,
		GithubEventName:        *consts.NewEventName(eventName),
		GithubWorkSpace:        workSpace,
		GithubBranch:           branch,
		GithubRunId:            runId,
		GithubServerUrl:        severUrl,
		GithubActionsStatus:    *consts.NewJobStatus(status),
		GithubActionsCommitMsg: commitMsg,
		GitHubActionsPrTitle:   prTitle,
		GitHubActionsPrUrl:     prUrl,
	}
}
