package services

import (
	"context"
	"eletronicMall/config"
	"eletronicMall/dao"
	"eletronicMall/model"
	"eletronicMall/pkg/code"
	"eletronicMall/pkg/utils"
	"eletronicMall/serializer"
	"fmt"
	"mime/multipart"
	"strings"

	"gopkg.in/mail.v2"
)

type UserService struct {
	Nickname string `form:"nickname" json:"nickname"`
	UserName string `form:"username" json:"username"`
	PassWord string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 验证码的key, 现阶段只要求前端验证
}

// SendEmailService 结构体用于接收发送电子邮件的请求参数
type SendEmailService struct {
	Email         string `json:"email" form:"email"`                   // Email 地址
	Password      string `json:"password" form:"password"`             // 密码
	OperationType uint   `json:"operation_type" form:"operation_type"` // 操作类型：1. 绑定邮箱，2. 解绑邮箱，3. 修改密码
}

// Register 用户注册
func (u *UserService) Register(ctx context.Context) serializer.Response {
	var user *model.User

	// 如果key 为空或者长度不为6, 则返回错误
	if u.Key == "" || len(u.Key) < 6 {
		return serializer.Response{
			Status:  code.InvalidParams,
			Message: code.GetMsg(code.InvalidParams),
			Error:   "密钥长度不足",
		}
	}

	// 初始金额 10000 --> 要转化成秘文存储, 对称加密
	utils.Encrypt.SetKey(u.Key)

	userDAO := dao.NewUserDAO(ctx)
	// 判断用户名是否存在
	_, exist, err := userDAO.CheckUserExist(u.UserName)

	// 如果在查询过程中出现错误, 则返回错误
	if err != nil {
		return serializer.Response{
			Status:  code.Error,
			Message: code.GetMsg(code.Error),
			Error:   err.Error(),
		}
	}

	// 如果用户名存在, 则返回错误
	if exist {
		return serializer.Response{
			Status:  code.UserExist,
			Message: code.GetMsg(code.UserExist),
			Error:   "用户名已存在",
		}
	}

	// 创建用户
	user = &model.User{
		Name:           u.UserName,
		PasswordDigest: u.PassWord,
		Email:          u.Nickname,
		Money:          utils.Encrypt.AesEncoding("10000"), // 对金额进行加密
		Status:         model.Active,
	}

	// 对密码进行加密
	if err := user.SetPassWord(u.PassWord); err != nil {
		return serializer.Response{
			Status:  code.Error,
			Message: code.GetMsg(code.Error),
			Error:   err.Error(),
		}
	}

	// 创建用户
	if err := userDAO.CreateUser(user); err != nil {
		return serializer.Response{
			Status:  code.Error,
			Message: code.GetMsg(code.Error),
			Error:   err.Error(),
		}
	}

	// 返回成功信息
	return serializer.Response{
		Status:  code.Success,
		Message: code.GetMsg(code.Success),
	}
}

// Login 用户登录
func (u *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User

	userdao := dao.NewUserDAO(ctx)
	user, exist, err := userdao.CheckUserExist(u.UserName)

	// 如果查询出错 or 用户不存在, 则返回错误
	if err != nil || !exist {
		return serializer.Response{
			Status:  code.Error,
			Message: code.GetMsg(code.Error),
			Error:   "用户名不存在",
		}
	}

	// 检查密码是否正确
	if !user.CheckPassword(u.PassWord) {
		return serializer.Response{
			Status:  code.Error,
			Message: code.GetMsg(code.Error),
			Error:   "密码错误",
		}
	}

	// 分发token
	// 因为http是无状态的, 意味着每次请求都不知道是谁, 所以需要token来标识用户
	token, err := utils.GenerateToken(user.ID, user.Name, 0)
	if err != nil {
		return serializer.Response{
			Status:  code.Error,
			Message: code.GetMsg(code.Error),
			Error:   err.Error(),
		}
	}

	return serializer.Response{
		Status:  code.Success,
		Message: code.GetMsg(code.Success),
		Data:    serializer.TokenData{Token: token, User: serializer.BuildUser(user)}, // 转化成前端所需要的数据类型
	}
}

// Update 用户更新
func (u *UserService) Update(ctx context.Context, uid uint) serializer.Response {

	userdao := dao.NewUserDAO(ctx)

	// 获取这个用户, 似乎在创建用户的时候没有指定id
	user, err := userdao.GetUserById(uid)
	if err != nil {
		return serializer.Response{
			Status:  code.Error,
			Message: code.GetMsg(code.Error),
			Error:   "用户不存在",
		}
	}

	// 修改昵称
	user.Name = u.Nickname
	err = userdao.UpdateById(uid, user)
	if err != nil {
		return serializer.Response{
			Status:  code.Error,
			Error:   err.Error(),
			Message: code.GetMsg(code.Error),
		}
	}

	return serializer.Response{
		Status:  code.Success,
		Message: code.GetMsg(code.Success),
		Data:    serializer.BuildUser(user),
	}
}

// UploadAvatar
func (u *UserService) UploadAvatar(
	ctx context.Context,
	uid uint,
	file multipart.File,
	fileSize int64) serializer.Response {

	userDao := dao.NewUserDAO(ctx)

	// 获取用户
	fmt.Println("获取用户")
	user, err := userDao.GetUserById(uid)
	if err != nil {
		return serializer.Response{
			Status:  code.Error,
			Error:   err.Error(),
			Message: code.GetMsg(code.Error),
		}
	}

	// 上传头像
	fmt.Println("上传头像")
	filepath, err := UploadAvatar(file, fileSize, uid, user.Name)
	if err != nil {
		return serializer.Response{
			Status:  code.Error,
			Error:   err.Error(),
			Message: code.GetMsg(code.Error),
		}
	}

	// 更新用户头像
	fmt.Println("更新头像")
	user.Avatar = filepath
	err = userDao.UpdateById(uid, user)
	if err != nil {
		return serializer.Response{
			Status:  code.Error,
			Error:   err.Error(),
			Message: code.GetMsg(code.Error),
		}
	}

	return serializer.Response{
		Status:  code.Success,
		Message: code.GetMsg(code.Success),
		Data:    serializer.BuildUser(user),
	}

}

func (s *SendEmailService) SendEmail(ctx context.Context, uid uint) serializer.Response {

	token, err := utils.GenerateEmailToken(uid, s.Email, s.Password, s.OperationType)
	if err != nil {
		return serializer.Response{
			Status:  code.Error,
			Error:   err.Error(),
			Message: code.GetMsg(code.Error),
		}
	}

	noticedao := dao.NewNoticeDAO(ctx)
	notice, err := noticedao.GetNoticeById(uid)
	if err != nil {
		return serializer.Response{
			Status:  code.Error,
			Error:   err.Error(),
			Message: code.GetMsg(code.Error),
		}
	}

	address := config.ValidEmail + token
	emailContent := notice.Text
	replacedContent := strings.Replace(emailContent, "Email", address, -1) // 替换掉emailContent中的Email

	msg := mail.NewMessage()
	msg.SetHeader("From", config.ValidEmail)
	msg.SetHeader("To", config.ValidEmail)
	msg.SetHeader("Subject", "绑定邮箱")
	msg.SetBody("text/html", replacedContent)
	deliver := mail.NewDialer(config.SmtpHost, 465, config.SmtpEmail, config.SmtpPass)
	deliver.StartTLSPolicy = mail.MandatoryStartTLS

	// 发送邮件
	// 如果发送邮件失败, 则返回错误
	if err := deliver.DialAndSend(msg); err != nil {
		return serializer.Response{
			Status:  code.Error,
			Error:   err.Error(),
			Message: code.GetMsg(code.Error),
		}
	}

	return serializer.Response{
		Status:  code.Success,
		Message: code.GetMsg(code.Success),
	}
}
