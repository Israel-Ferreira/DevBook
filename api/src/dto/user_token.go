package dto

type UserToken struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
}

func NewUserTokenDTO(userId uint, token string) *UserToken {
	return &UserToken{
		ID:    userId,
		Token: token,
	}
}
