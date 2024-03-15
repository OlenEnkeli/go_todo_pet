package schemas

import "github.com/OlenEnkeli/go_todo_pet/dto"

type UserBaseSchema struct {
	Name  string `json:"name" binding:"required"`
	Login string `json:"login" binding:"required"`
}

type UserCreateSchema struct {
	Password string `json:"password" binding:"required"`
	UserBaseSchema
}

type UserReturnSchema struct {
	Id int `json:"id"`
	UserBaseSchema
}

func (schema *UserCreateSchema) ToDTO() dto.User {
	return dto.User{
		Login:    schema.Login,
		Password: schema.Password,
		Name:     schema.Name,
	}
}

func (schema *UserReturnSchema) FromDTO(input dto.User) {
	schema.Id = input.Id
	schema.Login = input.Login
	schema.Name = input.Name
}

type LoginSchema struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
