package message

type ChatMessage struct {
	BaseMessage
	TargetUserid string `json:"target_userid"`
}
