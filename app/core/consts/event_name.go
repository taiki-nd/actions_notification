package consts

import "strings"

const (
	EventPush             = "push"
	EventPullRequest      = "pull_request"
	EventWorkflowDispatch = "workflow_dispatch"
)

type EventName struct {
	EventName string
}

func NewEventName(eventName string) *EventName {
	return &EventName{
		EventName: eventName,
	}
}

func (e *EventName) IsPush() bool {
	return e.EventName == EventPush
}

func (e *EventName) IsPullRequest() bool {
	return e.EventName == EventPullRequest
}

func (e *EventName) IsWorkflowDispatch() bool {
	return e.EventName == EventWorkflowDispatch
}

func (e *EventName) Value() string {
	return e.EventName
}

func (e *EventName) UPPERValue() string {
	return strings.ToUpper(e.EventName)
}
