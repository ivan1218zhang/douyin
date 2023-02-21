package mw

import (
	"github.com/golang-jwt/jwt/v4"
)

type myClaims struct {
	userId int64
	claim  jwt.RegisteredClaims
}

// JwtGetToken 根据userId通过jwt转换成token
func JwtGetToken(userId int64) (string, error) {
	return "", nil
}

// JwtGetUserId 根据token解密得到userId
func JwtGetUserId(token string) (int64, error) {
	return 0, nil
}

//func main() {
//	jwt_key := []byte("yingxiaozhu") // 加密key
//	// 加密一个token
//	claims := myClaims{
//		UserNmae: "颖小主",
//		StandardClaims: jwt.StandardClaims{
//			NotBefore: time.Now().Unix() - 60,      // 一分钟之前开始生效
//			ExpiresAt: time.Now().Unix() + 60*60*2, // 两个小时后失效
//			Issuer:    "签发人",                       // 签发人
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token_string, err := token.SignedString(jwt_key)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println(token_string) // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IumiluWwj-S4uyIsImV4cCI6MTY0OTk1NDkzOSwiaXNzIjoi562-5Y-R5Lq6IiwibmJmIjoxNjQ5OTQ3Njc5fQ.d8a24gGacP7Af_zy2NdUJvGO1-rHJENZXzV3dA_AESA
//
//	// 解密token
//	parseToken, err := jwt.ParseWithClaims(token_string, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtToken, nil
//	})
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	fmt.Println(parseToken.Claims.(*myClaims).UserNmae) // 颖小主
//
//}
