package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var TokenExpire = 3600
var RefreshTokenExpire = 7200
var JwtKey = "cat-disk-key"

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间
var CodeExpire = 500

var MailName = "xxx"
var MailPassword = "xxx"
