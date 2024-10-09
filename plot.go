package main

import (
	"image/color"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func initLinesPerAlgorithm() map[string]*plotter.Line {
	mapLines := make(map[string]*plotter.Line, 6)

	mapLines["Bubble Sort"] = &plotter.Line{}
	mapLines["Bubble Sort"].Color = color.RGBA{R: 255, G: 0, B: 100, A: 255}
	mapLines["Bubble Sort"].Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	mapLines["Selection Sort"] = &plotter.Line{}
	mapLines["Selection Sort"].Color = color.RGBA{R: 150, G: 50, B: 255, A: 255}
	mapLines["Selection Sort"].Dashes = []vg.Length{vg.Points(2), vg.Points(2)}

	mapLines["Insertion Sort"] = &plotter.Line{}
	mapLines["Insertion Sort"].Color = color.RGBA{R: 130, G: 255, B: 55, A: 255}
	mapLines["Insertion Sort"].Dashes = []vg.Length{vg.Points(13), vg.Points(13)}

	mapLines["Merge Sort"] = &plotter.Line{}
	mapLines["Merge Sort"].Color = color.RGBA{R: 255, G: 150, B: 170, A: 255}
	mapLines["Merge Sort"].Dashes = []vg.Length{vg.Points(10), vg.Points(10)}

	mapLines["Quick Sort"] = &plotter.Line{}
	mapLines["Quick Sort"].Color = color.RGBA{R: 150, G: 150, B: 225, A: 255}
	mapLines["Quick Sort"].Dashes = []vg.Length{vg.Points(15), vg.Points(15)}

	mapLines["Heap Sort"] = &plotter.Line{}
	mapLines["Heap Sort"].Color = color.RGBA{R: 50, G: 150, B: 255, A: 255}
	mapLines["Heap Sort"].Dashes = []vg.Length{vg.Points(20), vg.Points(20)}

	return mapLines
}



func plotGraphs(data map[string]map[string]map[string]map[string]interface{}, dist string, variable string, yLabel string, filename string) {
	f, err := os.Create("graphs/" + filename + ".png")
	if err != nil {
		log.Fatal(err)
	}
	p := plot.New()
	// p.Y.Scale = InvertedLogScale{}
	// p.Y.Tick.Marker = CustomLogTicks{}

	p.Title.Text = dist

	p.X.Label.Text = "Qtd. de itens"
	p.Y.Label.Text = yLabel

	mapStyles := initLinesPerAlgorithm()

	for k, v := range mapAlgorithms {
		XYs := plotter.XYs{}
		v = v
		for i := 0; i < 4; i++ {
			s, err := strconv.Atoi(listSizes[i])
			if err != nil {
				log.Fatal(err)
			}

			t := data[k][listSizes[i]][dist][variable]
			value := 0.0
			switch v := t.(type) {
				case int:
					value = float64(v)
				case float64:
					value = v
			}
			sFloat := float64(s)
			aux := plotter.XY{X: sFloat, Y: float64(value)}
			XYs = append(XYs, aux)
		}
		line, err := plotter.NewLine(XYs)
		if err != nil {
			log.Fatal(err)
		}
		line.Color = mapStyles[k].Color
		line.Dashes = mapStyles[k].Dashes

		p.Add(line)
		p.Legend.Add(k, line)
	}

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		log.Fatal(err)
	}

	_, err = wt.WriteTo(f)
	if err != nil {
		log.Fatal(err)
	}

	if err = f.Close(); err != nil {
		log.Fatal(err)
	}
}
