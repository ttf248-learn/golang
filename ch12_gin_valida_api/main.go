package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Username         string    `json:"username" binding:"required"`
	Email            string    `json:"email" binding:"required,email"`
	Age              int       `json:"age" binding:"required,gte=0,lte=130"`
	RegistrationDate time.Time `json:"registration_date" binding:"required,datetime"`
}

func main() {
	// 创建一个默认的路由引擎
	router := gin.Default()

	// 注册用户注册接口
	router.POST("/register", registerUser)

	// 运行HTTP服务器
	router.Run(":8080")
}

func registerUser(c *gin.Context) {
	// 绑定请求数据到User结构体
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 实例化校验器
	validate := validator.New()

	// 校验用户对象
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 如果校验通过，可以在这里保存用户信息到数据库或执行其他业务逻辑
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
}
