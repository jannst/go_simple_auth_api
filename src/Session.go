package src

type Session struct {
	SessionToken string `json:"token"`
	UserId uint32 `json:"id"`
}
