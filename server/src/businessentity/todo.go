package businessentity

/**
{“id”: string, “entityId”: string, “orgId”: string, “status”: string, “todoType”: string, “userId”: string}
*/

type TodoStatus string
type TodoType string

const (
	TodoStatusDone        = "DONE"
	TodoStatusPending     = "PENDING"
	TodoTypeChangeApprove = "CHANGE_APPROVE"
	TodoTypeFormSubmit    = "FORM_SUBMIT"
)

type Todo struct {
	ID       string     `json:"id"`
	EntityID string     `json:"entityId"`
	UserID   string     `json:"userId"`
	OrgID    string     `json:"orgId"`
	Status   TodoStatus `json:"status"`
	TodoType string     `json:"todoType"`
}
