package model

type User struct {
	ID   int    `json:"id"  gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}