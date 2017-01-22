package main

import (
	"qrtc/controller"
	r "qrtc/router"
)

func main() {

	//load configuration
	r.Router.LoadHTMLGlob("view/*")
	r.Router.Run("0.0.0.0:8080")

	defer controller.DB.Close()
}
