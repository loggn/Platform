package pkg

import (
	"fmt"

	//"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
}


var DB *gorm.DB

func InitDatabase() {
	var model User
	var err error
	dsn := "user=postgres password=123456 dbname=platformUser port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// DB, err = gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("连接数据库失败:", err)
	// 	return
	// }

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}

	// err = DB.AutoMigrate(model)
	// if err != nil {
	// 	fmt.Println("数据库迁移失败：", err)
	// 	return
	// }

	// fmt.Println("数据初始化成功")

	// 测试连接
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("获取 sqlDB 失败:", err)
		return
	}
	
	
	if err := sqlDB.Ping(); err != nil {
		fmt.Println("无法连接数据库:", err)
		return
	}
	fmt.Println("数据库连接成功")

	 // 使用 AutoMigrate 创建表
	 err = DB.AutoMigrate(model)
	 if err != nil {
		 fmt.Println("failed to migrate database:", err)
	 } else {
		 fmt.Println("database and table created successfully")
	 }
}
