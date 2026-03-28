package handler

import (
	"net/http"
	"strconv"

	"gin-demo/service"
	"gin-demo/utils"
	"github.com/gin-gonic/gin"
)
type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type UpdateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type LoginRequest struct {
	Name string `json:"name"`
}


func GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id必须是数字",
		})
		return
	}

	user, err := service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	if req.Name == "" || req.Age <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不合法"})
		return
	}

	user, err := service.CreateUser(req.Name, req.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": user.ID,
	})
}
func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	err := service.UpdateUser(id, req.Name, req.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}
func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数错误"})
		return
	}

	// 模拟用户（真实情况查数据库）
	if req.Name != "admin" {
		c.JSON(401, gin.H{"error": "用户不存在"})
		return
	}

	token, err := utils.GenerateToken(1)
	if err != nil {
		c.JSON(500, gin.H{"error": "生成token失败"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
func GetUserList(c *gin.Context) {
	// 1️⃣ 获取参数
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	// 2️⃣ 参数校验（很重要）
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	// 3️⃣ 调用 service
	list, total, err := service.GetUserList(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"error": "查询失败"})
		return
	}

	// 4️⃣ 返回结果
	c.JSON(200, gin.H{
		"list":  list,
		"total": total,
	})
}