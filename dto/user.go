package dto

type User struct {
	Id       uint
	Name     string
	Login    string
	Password string
}

type UserLogin struct {
	Login    string
	Password string
}
