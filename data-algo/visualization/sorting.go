package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/flan6/data-algo/algorithms/sorting"
)

func main() {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	length := 25
	selectionData := make([]float64, length)
	for i := range selectionData {
		selectionData[i] = float64(rand.Intn(length))
	}

	bubbleData := make([]float64, length)
	for i := range bubbleData {
		bubbleData[i] = float64(rand.Intn(length))
	}

	draw := func() {
		spline1 := widgets.NewSparkline()
		spline1.Data = selectionData
		spline1.LineColor = termui.ColorCyan
		spline1.Title = "SelectionSort"

		spline2 := widgets.NewSparkline()
		spline2.Data = bubbleData
		spline2.LineColor = termui.ColorMagenta
		spline2.Title = "BubbleSort"

		group := widgets.NewSparklineGroup(spline1, spline2)
		group.SetRect(0, 0, length, length)
		group.Title = "Sorting Visualization"

		bc := widgets.NewBarChart()
		bc.Data = selectionData

		// Create labels for each bar to display its value
		labels := make([]string, length)
		for i, v := range selectionData {
			labels[i] = fmt.Sprintf("%f", v)
		}
		bc.Labels = labels
		bc.Title = "Sorting Visualization"
		bc.SetRect(25, 0, 200, 25)
		bc.BarWidth = 2
		bc.LabelStyles = []termui.Style{termui.NewStyle(termui.ColorWhite)}
		bc.NumStyles = []termui.Style{termui.NewStyle(termui.ColorYellow)}

		termui.Render(group, bc)
	}

	draw()
	ticker := time.NewTicker(100 * time.Millisecond).C

	a := make(chan struct{})
	b := make(chan struct{})

	go sorting.SelectionSortVisual(selectionData, func() {
		a <- struct{}{}
	})

	go sorting.SelectionSortVisual(bubbleData, func() {
		b <- struct{}{}
	})

	termuiEvents := termui.PollEvents()
	for {
		select {
		case e := <-termuiEvents:
			if e.Type == termui.KeyboardEvent {
				return
			}

		case <-ticker:
			draw()
			<-a
			<-b
		}
	}
}
