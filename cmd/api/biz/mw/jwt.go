package mw

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

var MySecret = []byte("6379") // 定义secret，后面会用到

// JwtGetToken 根据userId通过jwt转换成token
func JwtGetToken(userId int64) (string, error) {
	// 加密一个token
	claims := &jwt.RegisteredClaims{
		Issuer:    "ivan",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
		IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
		NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		ID:        strconv.FormatInt(userId, 10),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(MySecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("6379"), nil // 这是我的secret
	}
}

// JwtGetUserId 根据token解密得到userId
func JwtGetUserId(tokens string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokens, &jwt.RegisteredClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return -1, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return -1, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return -1, errors.New("token not active yet")
			} else {
				return -1, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		id, err := strconv.ParseInt(claims.ID, 10, 64)
		if err != nil {
			return -1, err
		}
		return id, nil
	}
	return -1, errors.New("couldn't handle this token")
}
