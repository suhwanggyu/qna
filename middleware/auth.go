package middleware

import (
	"crypto/rsa"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/suhwanggyu/loginGo/controller"
)

func viperEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fail to read config file")
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal("Check .env file")
	}
	return value
}

// CheckAuth(pubkey *rsa.PublicKey) 미들웨어 함수이다.
// @dev TODO : 파라미터로 로그인 서버 주소를 넘겨 받도록 변경 예정
func CheckAuth(pubkey *rsa.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		if pubkey == nil {
			/* TODO : pass to host name */
			pubkey = controller.RequestPubKey()
		}
		var token controller.TokenExpired
		json.NewDecoder(c.Request.Body).Decode(&token)
		check := controller.CheckTokenExpired(*pubkey, token)
		if check == false {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusOK)
		c.Next()
	}
}
