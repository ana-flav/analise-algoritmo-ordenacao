package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

func plotGraphs(data map[string]map[string]map[string]map[string]interface{}) {
	// Plotando para Selection Sort
	XYs := plotter.XYs{}

	for i := range listSizes {
		s, err := strconv.Atoi(listSizes[i])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("plotando", listSizes[i])
		t := data["Merge Sort"][listSizes[i]]["Ordem decrescente"]["trocas"]
		v := t.(int)
		vFloat := float64(v)
		sFloat := float64(s)
		aux := plotter.XY{X: sFloat, Y: vFloat}
		XYs = append(XYs, aux)
	}

	plotLineGraph("out.png", "Qtd. Itens", "Qtd. Trocas", XYs)
}

func plotLineGraph(filename string, xLabel string, yLabel string, points plotter.XYs) {
	f, err := os.Create("graphs/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	p := plot.New()
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel
	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatal(err)
	}

	line.Color = color.RGBA{R: 255, A: 255}
	p.Add(line)

	sc, err := plotter.NewScatter(points)
	if err != nil {
		log.Fatal(err)
	}

	sc.Shape = draw.CircleGlyph{}
	sc.Color = color.RGBA64{R: 255, G: 125, B: 100}
	p.Add(sc)

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
