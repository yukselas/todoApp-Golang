package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var port string = DEFAULT_PORT

var server *http.Server

func InitServer() error {
	engine := gin.Default()

	engine.GET("/", getRoot)
	engine.GET("/notes", getNotes)
	engine.GET("/notes/:id", getNoteByID)
	engine.POST("/notes", postNote)
	engine.DELETE("/notes/:id", deleteNote)

	server = &http.Server{
		Addr:    ":" + port,
		Handler: engine,
	}
	return nil
}

func ReplySuccess(c *gin.Context, data interface{}) {
	yanit := gin.H{
		"data":      data,
		"status":    "success",
		"timestamp": time.Now().Unix(),
	}
	c.JSON(http.StatusOK, yanit)
}

func ReplyFail(c *gin.Context, data interface{}) {

	yanit := gin.H{
		"data":   data,
		"status": "fail",
	}
	c.JSON(http.StatusBadRequest, yanit)
}

func RunServer() error {
	return server.ListenAndServe()
}
