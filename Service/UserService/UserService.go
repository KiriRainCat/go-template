package UserService

import (
	"go-template/Common/DB/MySQL"
	"go-template/Model/Entity"
)

func GetUserByID(id string) (res Entity.User, err error) {
	err = MySQL.DB.First(&res, "id = ?", id).Error
	return
}

func AddUser(user Entity.User) (id int, err error) {
	err = MySQL.DB.Create(&user).Error
	id = user.ID
	return
}
