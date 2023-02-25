package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func main() {
	key := []byte("my_secret_key")


	// 生成jwt
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1234567890"
	claims["name"] = "jerry bob"
	claims["iat"] = time.Now().Unix()   //令牌的签发时间
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()   //令牌的有效期

	// 签名
	tokenString,err := token.SignedString(key)
	if err != nil {
		fmt.Println("Failed to generate token")
		return
	}

	fmt.Println(tokenString)

	// 解析jwt
	token1,err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC);!ok {
			return nil,fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// 返回秘钥
		return key,nil
	})
	if err != nil {
		fmt.Println("Failed to parse token")
		return
	}

	// 验证载荷
	if claims,ok := token1.Claims.(jwt.MapClaims);ok && token1.Valid{
		sub := claims["sub"].(string)
		name := claims["name"].(string)
		iat := int64(claims["iat"].(float64))
		exp := int64(claims["exp"].(float64))

		fmt.Printf("Subject: %s\nName: %s\nIssued at: %s\nExpiration time: %s\n", sub, name, time.Unix(iat, 0), time.Unix(exp, 0))
	}else {
		fmt.Println("Invalid token")
	}


}
