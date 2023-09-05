package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt: json web token
var JwtSecret = []byte("eletronicMall")

type Claims struct {
	UserId             uint   `json:"user_id"`
	Username           string `json:"username"`
	Authority          int    `json:"authority"`
	jwt.StandardClaims        // jwt的标准字段
}

// EmailClaims 结构体用于定义电子邮件认证的 JWT（JSON Web Token）声明
type EmailClaims struct {
	UserID             uint   `json:"user_id"`        // 用户ID
	Email              string `json:"email"`          // 邮箱地址
	Password           string `json:"password"`       // 密码
	OperationType      uint   `json:"operation_type"` // 操作类型, 0 为绑定邮箱, 1 为解绑邮箱, 2 为修改密码
	jwt.StandardClaims        // JWT 标准声明，包含了标准的 JWT 字段（例如过期时间、发行人等）
}

// GenerateToken 生成token
func GenerateToken(userId uint, username string, authority int) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		UserId:    userId,
		Username:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			Issuer:    "eletronicMall",   // 签发人
		},
	}

	// 1. 创建一个新的令牌对象，指定签名方法和标准声明
	// 使用  HMAC-SHA256 算法生成签名，使用对称密钥，即发送方和接收方共享同一个密钥来生成和验证认证码。
	// 速度比较快，适合对大块文本进行签名和验证，安全性高。
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 2. 使用指定的密钥生成签名字符串
	token, err := tokenClaims.SignedString(JwtSecret)

	return token, err
}

// ParseToken 解析token
//
// 原理是使用相同的密钥来验证令牌的签名是否有效
func ParseToken(token string) (*Claims, error) {

	// 1. 解析token
	// 在验证令牌时，需要提供用于验证签名的密钥，
	// 但不应将密钥硬编码在代码中或在公开的地方存储。回调函数提供了一种机制，可以在需要时从安全的位置获取密钥，而不必将其明文存储在代码中。
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 2. 返回密钥
		return JwtSecret, nil
	})

	// 3. 判断token是否有效
	if tokenClaims != nil {
		// 4. 判断tokenClaims中的Claims是否有效
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GenerateEmailToken 生成电子邮件认证的 JWT（JSON Web Token）
func GenerateEmailToken(userID uint, email string, password string, operationType uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := EmailClaims{
		UserID:        userID,
		Email:         email,
		Password:      password,
		OperationType: operationType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "eletronicMall", // 替换成实际的签发人
		},
	}

	// 创建一个新的令牌对象，指定签名方法和声明
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用指定的密钥生成签名字符串
	token, err := tokenClaims.SignedString(JwtSecret)

	return token, err
}

// ParseEmailToken 解析电子邮件认证的 JWT（JSON Web Token）
//
// 原理是使用相同的密钥来验证令牌的签名是否有效
func ParseEmailToken(token string) (*EmailClaims, error) {
	// 1. 解析 Token
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 2. 返回密钥
		return JwtSecret, nil
	})

	// 3. 判断 Token 是否有效
	if tokenClaims != nil {
		// 4. 判断 TokenClaims 中的 Claims 是否有效
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
