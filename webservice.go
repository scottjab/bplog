package bplog

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (b *BPLog) ServeHttp() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "pong",
			"time_stamp": time.Now(),
		})
	})
	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("ui/build"))))
	r.Run()
}
