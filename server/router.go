package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"leopard/api"
	"leopard/middleware"
	"log"
	"time"

	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)
		//TODO:https://github.com/yangyongzhen/gin-websocket

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)

		}
	}
	return r
}

// websocket连接处理
func WsConnHandle(conn *websocket.Conn) {
	for {
		fmt.Println("6666")
		var msg string
		if err := websocket.Message.Receive(conn, &msg); err != nil {
			log.Println(err)
			return
		}

		log.Printf("recv: %v", msg)

		data := []byte(time.Now().Format(time.RFC3339))
		if _, err := conn.Write(data); err != nil {
			log.Println(err)
			return
		}
	}
}
