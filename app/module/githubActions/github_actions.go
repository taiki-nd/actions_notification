package githubActions

type GithubActionsClient struct {
	GithubRepository string
	GithubSha        string
	GithubRef        string
	GithubActor      string
	GithubWorkflow   string
	GithubEventName  string
	GithubWorkSpace  string
	GithubBranch     string
	GithubRunId      string
	GithubServerUrl  string
}

func NewGithubActions(repo, sha, ref, actor, workflow, eventName, workSpace, branch, runId, severUrl string) *GithubActionsClient {
	return &GithubActionsClient{
		GithubRepository: repo,
		GithubSha:        sha,
		GithubRef:        ref,
		GithubActor:      actor,
		GithubWorkflow:   workflow,
		GithubEventName:  eventName,
		GithubWorkSpace:  workSpace,
		GithubBranch:     branch,
		GithubRunId:      runId,
		GithubServerUrl:  severUrl,
	}
}
