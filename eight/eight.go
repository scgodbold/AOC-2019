package eight

import (
	"fmt"
)

type layer struct {
	Pixels [][]int
	Counts map[int]int
}

func newLayer(h int, w int) layer {
	var l layer
	l.Counts = make(map[int]int)
	l.Pixels = make([][]int, h)
	for i := 0; i < h; i++ {
		l.Pixels[i] = make([]int, w)
	}

	return l
}

func buildLayers(input []string, h int, w int) []layer {
	layers := []layer{
		newLayer(h, w),
	}

	pointer := 0
	curHeight := 0
	curWidth := 0
	for i, v := range input[0] {
		val := int(v - '0')
		layers[pointer].Pixels[curHeight][curWidth] = val
		if _, ok := layers[pointer].Counts[val]; !ok {
			layers[pointer].Counts[val] = 1
		} else {
			layers[pointer].Counts[val] += 1
		}

		if i == len(input[0])-1 {
			break
		}

		curWidth += 1
		if curWidth >= w {
			curWidth = 0
			curHeight += 1
		}

		if curHeight >= h {
			curHeight = 0
			pointer += 1
			layers = append(layers, newLayer(h, w))
		}
	}

	return layers
}

func processFewestZeros(layers []layer) {
	pointer := 0
	for i, _ := range layers {
		if i == 0 {
			continue
		}

		if layers[i].Counts[0] < layers[pointer].Counts[0] {
			pointer = i
		}
	}

	fmt.Println(layers[pointer].Counts[1] * layers[pointer].Counts[2])
}

func buildImage(layers []layer, h int, w int) layer {
	layer := newLayer(h, w)
	for i := len(layers) - 1; i >= 0; i-- {
		for height := 0; height < h; height++ {
			for width := 0; width < w; width++ {
				val := layers[i].Pixels[height][width]
				if i == len(layers)-1 {
					layer.Pixels[height][width] = val
				} else if val != 2 {
					layer.Pixels[height][width] = val
				}
			}
		}
	}
	return layer
}

func printImage(image layer) {
	for _, v := range image.Pixels {
		fmt.Println(v)
	}
}

func Run(input []string) {
	height := 6
	width := 25
	layers := buildLayers(input, height, width)
	processFewestZeros(layers)

	image := buildImage(layers, height, width)
	printImage(image)
}
