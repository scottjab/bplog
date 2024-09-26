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
	r.GET("/api/bp", b.ServeBPData)
	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("ui/build"))))
	r.Run()
}

type BPChartData struct {
	Dates     []time.Time `json:"dates"`
	Systolic  []int       `json:"systolic"`
	Diastolic []int       `json:"diastolic"`
	Pulse     []int       `json:"pulse"`
}

func (b *BPLog) ServeBPData(ctx *gin.Context) {
	readings := []BloodPressure{}
	r := b.db.Find(&readings)

	chart := &BPChartData{
		Dates:     make([]time.Time, r.RowsAffected),
		Systolic:  make([]int, r.RowsAffected),
		Diastolic: make([]int, r.RowsAffected),
		Pulse:     make([]int, r.RowsAffected),
	}

	for i, reading := range readings {
		chart.Dates[i] = reading.Date
		chart.Systolic[i] = reading.Systolic
		chart.Diastolic[i] = reading.Diastolic
		chart.Pulse[i] = reading.Pulse
	}
	ctx.JSON(http.StatusOK, chart)
}
