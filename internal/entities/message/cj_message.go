package message

type ChatMessage struct {
	BaseChatMessage
	TargetUserid string `json:"target_userid"`
}
