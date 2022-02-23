package dto

type FollowerDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
}
