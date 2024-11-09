package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	功能：注册接口
	描述：注册新用户，成功后返回用户信息和认证 Token
	创建时间：2024-11-6
*/
type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func RegisterHandler(ctx *gin.Context) {
	var request registerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing parameters",
		})
		return
	}

	user := User{}
	if err := DB.Where("username = ?", request.Username).First(&user).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	if err := DB.Where("email = ?", request.Email).First(&user).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	hashedPassword, err := HashPassword(request.Password)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to hash password",
        })
        return
    }

	user = User{
		Username: request.Username,
		Password: hashedPassword,
		Email:    request.Email,
	}

	if err := DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "User fail to register",
		})
		fmt.Println("User fail to register", err)
		return
	}

	token, err := GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "token create fail",
		})
		fmt.Println("token create fail", err)
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"userId": user.ID,
		"username": user.Username,
		"email": user.Email,
		"token": token,
	})	
}


/*
	功能：登录接口
	描述：用户登录，成功后返回用户信息和认证 Token。
	创建时间：2024年11月6日
*/
type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(ctx *gin.Context) {
	var request loginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing parameters",
		})
		return
	}

	var user User
	if err := DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	if isCorrect := CheckPasswordHash(request.Password, user.Password); isCorrect {
		token, err := GenerateToken(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "token create fail",
			})
			fmt.Println("token create fail", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"userId": user.ID,
			"username": user.Username,
			"email": user.Email,
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
	}
}