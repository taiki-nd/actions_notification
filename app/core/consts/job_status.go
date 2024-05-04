package consts

import "strings"

const (
	JobStatusSuccess = "success"
	JobStatusFailure = "failure"
	JobStatusCancel  = "cancelled"
)

type JobStatus struct {
	Status string
}

func NewJobStatus(status string) *JobStatus {
	return &JobStatus{
		Status: status,
	}
}

func (j *JobStatus) IsSuccess() bool {
	return j.Status == JobStatusSuccess
}

func (j *JobStatus) IsFailure() bool {
	return j.Status == JobStatusFailure
}

func (j *JobStatus) IsCancelled() bool {
	return j.Status == JobStatusCancel
}

func (j *JobStatus) Value() string {
	return j.Status
}

func (j *JobStatus) UPPERValue() string {
	return strings.ToUpper(j.Status)
}

func (j *JobStatus) IconValue() string {
	if j.IsSuccess() {
		return ":white_check_mark:"
	}
	if j.IsFailure() {
		return ":x:"
	}
	if j.IsCancelled() {
		return ":warning:"
	}
	return ""
}
