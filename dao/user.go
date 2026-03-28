package dao

import "gin-demo/model"

func GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := DB.First(&user, id).Error
	return &user, err
}

func CreateUser(user *model.User) error {
	return DB.Create(user).Error
}

func UpdateUser(id int, name string, age int) error {
	return DB.Model(&model.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name": name,
			"age":  age,
		}).Error
}

func DeleteUser(id int) error {
	return DB.Delete(&model.User{}, id).Error
}

func GetUserList(page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	DB.Model(&model.User{}).Count(&total)

	offset := (page - 1) * pageSize
	err := DB.Limit(pageSize).Offset(offset).Find(&users).Error

	return users, total, err
}
