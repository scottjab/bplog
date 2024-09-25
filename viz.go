package bplog

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
)

func (b *BPLog) Plot() {
	readings := []BloodPressure{}

	b.db.Find(&readings)

	systolic := []float64{}
	diastolic := []float64{}
	pulse := []float64{}
	for _, reading := range readings {
		systolic = append(systolic, float64(reading.Systolic))
		diastolic = append(diastolic, float64(reading.Diastolic))
		pulse = append(pulse, float64(reading.Pulse))
	}
	data := [][]float64{
		systolic,
		diastolic,
		//pulse,
	}
	graph := asciigraph.PlotMany(data, asciigraph.Precision(0), asciigraph.Height(50), asciigraph.Width(200), asciigraph.SeriesColors(
		asciigraph.Green,
		asciigraph.Red,
		//asciigraph.Blue,
	), asciigraph.SeriesLegends("systolic", "diastolic"))
	fmt.Println(graph)

}
