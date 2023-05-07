package user

type IUser interface {
	GetUserId() string
	SetLastPing(time int64)
	GetLastPing() int64
}
type BaseUser struct {
	UserId string `json:"userid"`
}
