package githubActions

import "github.com/taiki-nd/actions_notification/app/core/consts"

type GithubActionsClient struct {
	GithubRepository string
	GithubSha        string
	GithubRef        string
	GithubActor      string
	GithubWorkflow   string
	GithubEventName  consts.EventName
	GithubWorkSpace  string
	GithubBranch     string
	GithubRunId      string
	GithubServerUrl  string
	GithubJobStatus  consts.JobStatus
	GithubCommitMsg  string
	GitHubPrTitle    string
	GitHubPrUrl      string
	GitHubBaseRef    string
	GitHubHeadRef    string
}

func NewGithubActions(repo, sha, ref, actor, workflow, eventName, workSpace, branch, runId, severUrl, status, commitMsg, prTitle, prUrl, baseRef, headRef string) *GithubActionsClient {
	return &GithubActionsClient{
		GithubRepository: repo,
		GithubSha:        sha,
		GithubRef:        ref,
		GithubActor:      actor,
		GithubWorkflow:   workflow,
		GithubEventName:  *consts.NewEventName(eventName),
		GithubWorkSpace:  workSpace,
		GithubBranch:     branch,
		GithubRunId:      runId,
		GithubServerUrl:  severUrl,
		GithubJobStatus:  *consts.NewJobStatus(status),
		GithubCommitMsg:  commitMsg,
		GitHubPrTitle:    prTitle,
		GitHubPrUrl:      prUrl,
		GitHubBaseRef:    baseRef,
		GitHubHeadRef:    headRef,
	}
}
