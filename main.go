package main

import (
	"crypto/rsa"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suhwanggyu/loginGo/controller"
	"github.com/suhwanggyu/loginGo/key"
)

var pubkey *rsa.PublicKey = nil

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if pubkey == nil {
			pubkey = RequestPubKey()
		}
		token := controller.TokenExpired{}
		check := key.CheckTokenExpired(*pubkey, token)
		if check == false{
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORS())
	router.Use(CheckAuth())
	router.Run()
}
