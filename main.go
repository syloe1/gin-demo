package main

import (
	"gin-demo/dao"
	"gin-demo/model"
	"gin-demo/router"
)

func main() {
	dao.InitDB()
	dao.InitRedis()
	dao.DB.AutoMigrate(&model.User{})
	r := router.SetupRouter()
	r.Run(":10000")
}
