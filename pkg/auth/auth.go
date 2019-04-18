/**
 * 功能描述: 用户认证
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package auth

import "golang.org/x/crypto/bcrypt"

// 将文本进行加密
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// 比较密文密码与明文密码是否一致
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
