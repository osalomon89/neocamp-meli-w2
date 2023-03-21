package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	const port string = ":9000"
	r := gin.Default()

	r.GET("/ping", ping)

	log.Println("Server listening on port", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
}

type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ResponseInfo{
		Error: false,
		Data:  "pong",
	})
}
