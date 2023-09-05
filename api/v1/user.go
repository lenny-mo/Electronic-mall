package v1

import (
	"eletronicMall/pkg/utils"
	"eletronicMall/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 注册 controller
func Register(c *gin.Context) {
	var userRegister services.UserService
	//
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)

	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// Login 用户登录
func Login(c *gin.Context) {
	var userLogin services.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// UserUpdate 用户更新
func UserUpdate(c *gin.Context) {
	var userUpdate services.UserService

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))

	// 从请求体中获取参数
	if err := c.ShouldBind(&userUpdate); err == nil {
		// 根据claims中的用户id, 更新用户信息
		res := userUpdate.Update(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// UserUploadAvatar 用户上传头像
func UserUploadAvatar(c *gin.Context) {
	// https://images6.alphacoders.com/132/1322714.jpeg

	file, fileHeader, err := c.Request.FormFile("avatar")
	if err != nil {
		fmt.Println("文件上传失败第一步", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件上传失败",
		})
		return
	}

	fileSize := fileHeader.Size

	var upploadAvatarService services.UserService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&upploadAvatarService); err == nil { // xxx: 这一步有问题
		fmt.Println("进入上传头像")
		res := upploadAvatarService.UploadAvatar(c.Request.Context(), claims.UserId, file, fileSize)
		c.JSON(http.StatusOK, res)
	} else {
		fmt.Println("进入上传头像失败")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件上传失败",
		})
	}
}

// SendEmail 发送邮件
func SendEmail(c *gin.Context) {
	var sendEmailService services.SendEmailService

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))	

	if err := c.ShouldBind(&sendEmailService); err == nil {
		res := sendEmailService.SendEmail(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}

	
}
