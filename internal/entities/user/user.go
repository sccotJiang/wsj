package user

type IUser interface {
	GetUserId() string
	SetLastPing(time int64)
	GetLastPing() int64
}
