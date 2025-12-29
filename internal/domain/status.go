package domain

// TalkStatus represents the status of a talk submission.
type TalkStatus string

const (
	StatusSubmitted TalkStatus = "SUBMITTED"
	StatusApproved  TalkStatus = "APPROVED"
	StatusRejected  TalkStatus = "REJECTED"
	StatusDraft     TalkStatus = "DRAFT"
	StatusWithdrawn TalkStatus = "WITHDRAWN"
)

// IsPublic returns true if the talk status indicates it should be publicly visible.
// Only approved talks should be visible in the public index.
func (t TalkStatus) IsPublic() bool {
	return t == StatusApproved
}
