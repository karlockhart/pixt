package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/karlockhart/pixt/internal/app/geosim"
)

func run() {

	nm := geosim.NewNodeMesh(1024, 768)

	err := nm.SetHeight(1000, 400, 100)
	if err != nil {
		fmt.Println(err)
	}
	err = nm.SetHeight(400, 600, 120)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nm.MaxSortedHeight.Height)

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
