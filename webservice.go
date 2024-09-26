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
	Labels   []string  `json:"labels"`
	Datasets []DataSet `json:"datasets"`
}
type DataSet struct {
	Name      string `json:"name"`
	ChartType string `json:"chartType"`
	Values    []int  `json:"values"`
}

func (b *BPLog) ServeBPData(ctx *gin.Context) {
	readings := []BloodPressure{}
	r := b.db.Find(&readings)
	chart := BPChartData{
		Labels: make([]string, r.RowsAffected),
	}
	systolic := DataSet{
		Name:      "Systolic",
		ChartType: "line",
		Values:    make([]int, r.RowsAffected),
	}

	diastolic := DataSet{
		Name:      "Diastolic",
		ChartType: "line",
		Values:    make([]int, r.RowsAffected),
	}
	pulse := DataSet{
		Name:      "Pulse",
		ChartType: "line",
		Values:    make([]int, r.RowsAffected),
	}
	for i, reading := range readings {
		chart.Labels[i] = reading.Date.String()
		systolic.Values[i] = reading.Systolic
		diastolic.Values[i] = reading.Diastolic
		pulse.Values[i] = reading.Pulse
	}
	chart.Datasets = []DataSet{systolic, diastolic, pulse}
	ctx.JSON(http.StatusOK, chart)
}
