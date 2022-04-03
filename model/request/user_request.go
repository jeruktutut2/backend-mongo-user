package request

type UserRequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
