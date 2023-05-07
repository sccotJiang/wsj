package message

const ( //消息大类
	TypeNormal = "normal"
)

const ( //消息子类
	SubTypePublicChat  = "publicchat"
	SubTypePrivateChat = "privateChat"
)

type BaseMessage struct {
	Type      string `json:"type"`
	Namespace string `json:"namespace,omitempty"`
	MsgBody   string `json:"msg_body"`
}
