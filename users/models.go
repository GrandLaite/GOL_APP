package users

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsPremium bool   `json:"is_premium"`
}
