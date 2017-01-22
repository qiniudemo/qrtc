package router

import (
	"github.com/gin-gonic/gin"

	"qrtc/controller"
)

var (
	Router = gin.Default()
)

func init() {

	//routers
	Router.GET("/", controller.Index)

	Router.GET("/create", controller.GetCreateRoom)
	Router.POST("/create", controller.PostCreateRoom)

	Router.GET("/room", controller.GetCheckRoom)
	Router.POST("/room", controller.PostCheckRoom)
	Router.GET("/rooms", controller.GetAllRooms)

	Router.GET("/delete", controller.GetDeleteRoom)
	Router.POST("/delete", controller.PostDeleteRoom)

	Router.GET("/token", controller.GetToken)
	Router.POST("/token", controller.TokenGenerator)
	Router.POST("/room/:room_name/user/:user_id/token", controller.TokenAPI)

	Router.GET("/login", controller.GetLogin)
	Router.POST("/login", controller.PostLogin)
	
	Router.GET("/stream",controller.GetPublishUrl)
	Router.POST("/stream",controller.PostPublisUrl)
	Router.POST("/stream/:room_name",controller.GenPubUrl)
	
	Router.GET("/stream/:room_name/play",controller.GenPlayUrl)
	

	Router.StaticFile("/bootstrap.min.css", "./view/bootstrap.min.css")
	Router.StaticFile("/jquery.min.js", "./view/jquery.min.js")

}
