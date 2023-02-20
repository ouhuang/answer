package core

import "2333/answer/initialize"

func RunServer() {
	Router := initialize.Routers()
	// Router.Static("/form-generator", "./resource/page")

	Router.Run(":8080")
}
