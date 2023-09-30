package model

import "github.com/nghuyentrangvt98/banhmioi/ent"

type UserCreate struct {
	Username string `json:"username,omitempty"`
	Fullname string `json:"fullname,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Dob      string `json:"dob,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type UserLogin struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Fullname string `json:"-"`
	Gender   string `json:"-"`
	Dob      string `json:"-"`
	Password string `json:"-"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	JWT      string `json:"jwt,omitempty"`
}

func FromSchemaUser(user *ent.User, jwt string) *UserResponse {
	return &UserResponse{
		Id:       user.ID,
		Username: user.Username,
		JWT:      jwt,
	}
}
