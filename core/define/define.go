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
