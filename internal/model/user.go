package model

type UserRegisterInput struct {
	UserName string
	NickName string
	Password string
}

type UserUpdateInput struct {
	UserName string
	NickName string
	Password string
}
