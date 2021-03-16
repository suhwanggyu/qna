package main

import (
	"crypto/rsa"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/suhwanggyu/qna/db"
)

var pubkey *rsa.PublicKey = nil
type routerType struct{
	*gin.Engine
}

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

func (router routerType) RoutesSetting() {
	router.GET("/thread", func(c *gin.Context) {
		fmt.Println("TODO")
	})
	router.POST("/thread", func(c *gin.Context) {
		fmt.Println("TODO")
	})
}

func main() {
	router := &routerType{gin.Default()}
	router.Use(CORS())
	//router.Use(middleware.CheckAuth(pubkey))
	db.M()
	router.RoutesSetting()
	router.Run(":3000")
}