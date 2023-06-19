package UserService

import (
	"go-template/Common/DB/MySQL"
	"go-template/Model/Entity"
)

func GetUserByID(id string) Entity.User {
	var user Entity.User
	MySQL.DB.First(&user, id)
	return user
}

func AddUser(user Entity.User) error {
	res := MySQL.DB.Create(&user)
	return res.Error
}
