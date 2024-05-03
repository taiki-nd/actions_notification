package githubActions

import "actionsnotification/app/core/consts"

type GithubActionsClient struct {
	GithubRepository       string
	GithubSha              string
	GithubRef              string
	GithubActor            string
	GithubWorkflow         string
	GithubEventName        string
	GithubWorkSpace        string
	GithubBranch           string
	GithubRunId            string
	GithubServerUrl        string
	GithubActionsStatus    consts.JobStatus
	GithubActionsCommitMsg string
}

func NewGithubActions(repo, sha, ref, actor, workflow, eventName, workSpace, branch, runId, severUrl, status, commitMsg string) *GithubActionsClient {
	jobStatus := consts.NewJobStatus(status)
	return &GithubActionsClient{
		GithubRepository:       repo,
		GithubSha:              sha,
		GithubRef:              ref,
		GithubActor:            actor,
		GithubWorkflow:         workflow,
		GithubEventName:        eventName,
		GithubWorkSpace:        workSpace,
		GithubBranch:           branch,
		GithubRunId:            runId,
		GithubServerUrl:        severUrl,
		GithubActionsStatus:    *jobStatus,
		GithubActionsCommitMsg: commitMsg,
	}
}
