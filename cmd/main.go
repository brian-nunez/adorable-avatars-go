package main

import (
	"github.com/brian-nunez/adorable-avatars-go/cmd/avatars"
	"github.com/brian-nunez/adorable-avatars-go/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := avatars.GetFaceOptionsWithoutCache()
	if err != nil {
		panic("Could not load required resources.\nExiting...")
	}

	r := gin.Default()

	r.GET("/list", handlers.ListGETHandler)
	r.GET("/random", handlers.RandomGETHandler)

	r.Run()
}
