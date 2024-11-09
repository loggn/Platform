package main

import (
	"log"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang/pkg"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置示例变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}
	

func main() {
	/*
	功能：初始化数据库
	描述：将User模型传入初始化数据库API，完成数据库的初始化
	创建时间：2024年11月6日
	*/
	pkg.InitDatabase()
	sqlDB, err := pkg.DB.DB()
	if err != nil {
		fmt.Println("获取 sqlDB 失败:", err)
		return
	}
	defer sqlDB.Close()

	r := gin.Default()

	r.Use(Logger())

	/**
	功能：注册接口
	描述：注册新用户，成功后返回用户信息和认证 Token
	创建时间：2024年11月6日
	**/
	r.POST("/api/register", pkg.RegisterHandler)

	/*
	功能：登录接口
	描述：用户登录，成功后返回用户信息和认证 Token。
	创建时间：2024年11月6日
	*/
	r.POST("/api/login", pkg.LoginHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "World"
		}
		c.JSON(200, gin.H{
			"message": "Hello " + name,
		})
	})

	r.Run(":5488") // 默认在 0.0.0.0:8080 启动服务
}
