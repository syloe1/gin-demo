package service

import (
	"encoding/json"
	"fmt"
	"gin-demo/dao"
	"gin-demo/model"
)

func GetUser(id int) (*model.User, error) {
	key := fmt.Sprintf("user:%d", id)

	// 1️⃣ 先查 Redis
	val, err := dao.RDB.Get(dao.Ctx, key).Result()
	if err == nil {
		var user model.User
		json.Unmarshal([]byte(val), &user)
		return &user, nil
	}

	// 2️⃣ Redis 没有 → 查 MySQL
	user, err := dao.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// 3️⃣ 写入 Redis（缓存）
	data, _ := json.Marshal(user)
	dao.RDB.Set(dao.Ctx, key, data, 0) // 0 = 不过期

	return user, nil
}

func CreateUser(name string, age int) (*model.User, error) {
	user := &model.User{
		Name: name,
		Age:  age,
	}
	err := dao.CreateUser(user)
	return user, err
}

func UpdateUser(id int, name string, age int) error {
	err := dao.UpdateUser(id, name, age)
	if err == nil {
		dao.RDB.Del(dao.Ctx, fmt.Sprintf("user:%d", id))
	}
	return err
}

func DeleteUser(id int) error {
	err := dao.DeleteUser(id)
	if err == nil {
		dao.RDB.Del(dao.Ctx, fmt.Sprintf("user:%d", id))
	}
	return err
}

func GetUserList(page, pageSize int) ([]model.User, int64, error) {
	return dao.GetUserList(page, pageSize)
}
