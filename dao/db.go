package dao

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client
var Ctx = context.Background()

func InitDB() {
	dsn := "root:123456@tcp(127.0.0.1:3307)/test?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
}

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码留空
		DB:       0,
	})

	if err := RDB.Ping(Ctx).Err(); err != nil {
		log.Fatal("Redis连接失败:", err)
	}
}
