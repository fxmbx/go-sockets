package main

import "github.com/gin-gonic/gin"

func main() {
	go h.run()
	router := gin.New()

	router.LoadHTMLFiles("index.htm")
	router.GET("/room/:roomId", func(ctx *gin.Context) {
		ctx.HTML(200, "index.htm", nil)
	})

	router.GET("/ws/:roomId", func(ctx *gin.Context) {
		roomId := ctx.Param("roomId")
		serveWs(ctx.Writer, ctx.Request, roomId)
	})

	router.Run("0.0.0.0:8085")
}
