package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/send", func(c *gin.Context) {
		send(c.Writer, c.Request)
	})
	engine.GET("/long_polling", func(c *gin.Context) {
		longPolling(c.Writer, c.Request)
	})
	
	engine.Run(":3000")
}

var msgCh = make(chan string)

func send(w http.ResponseWriter, r *http.Request) {
	msgCh <- "hello"

	w.Write([]byte("ok"))
}

func longPolling(w http.ResponseWriter, r *http.Request) {
	// msgChへ値が送信されるまで処理をブロック
	msg := <-msgCh

	w.Write([]byte(msg))
}
