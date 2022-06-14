package models

type UserAuth struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
}
