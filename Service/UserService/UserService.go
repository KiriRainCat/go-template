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

func UpdateUserByID(id string, user Entity.User) (affected int64, err error) {
	res := MySQL.DB.Model(&user).Where("id = ?", id).Updates(&user)
	affected = res.RowsAffected
	err = res.Error
	return
}

func DeleteUserByID(id string) (affected int64, err error) {
	res := MySQL.DB.Delete(&Entity.User{}, "id = ?", id)
	affected = res.RowsAffected
	err = res.Error
	return
}
