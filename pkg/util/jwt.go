package util

// JwtGetToken 根据userId通过jwt转换成token
func JwtGetToken(userId int64) (string, error) {
	return "", nil
}

// JwtGetUserId 根据token解密得到userId
func JwtGetUserId(token string) (int64, error) {
	return 0, nil
}
