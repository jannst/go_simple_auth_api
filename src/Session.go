package src

type Session struct {
	SessionToken string `json:"token"`
	UserId uint32 `json:"id"`
	UserRole string `json:"role"`
	UserName string `json:"name"`
	UserEmail string `json:"email"`
}
