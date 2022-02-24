package dto

type FollowingUserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
}
