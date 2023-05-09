package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// 生成随机盐
func generateSalt(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// 基于字符串进行摘要算法并加盐
func digestWithSalt(data string) (string, string) {
	salt := generateSalt(8)
	saltedData := salt + data
	hasher := md5.New()
	hasher.Write([]byte(saltedData))
	digest := hex.EncodeToString(hasher.Sum(nil))
	return digest, salt
}

func GenMD5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

// 校验摘要值是否匹配
func validateWithSalt(data string, salt string, digest string) bool {
	saltedData := salt + data
	hasher := md5.New()
	hasher.Write([]byte(saltedData))
	return digest == hex.EncodeToString(hasher.Sum(nil))
}

// 示例用法
func main() {
	data := "password123"
	digest, salt := digestWithSalt(data)
	fmt.Println("摘要值：", digest)
	fmt.Println("盐值：", salt)
	fmt.Println("是否匹配：", validateWithSalt(data, salt, digest))
}
