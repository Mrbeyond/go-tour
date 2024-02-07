package main

import (
	"math"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			switch ((x + y) / 2) % 3 {
			case 0:
				row[x] = uint8((x + y) / 2)
			case 1:
				row[x] = uint8(math.Pow(float64(x), float64(y)))
			default:
				row[x] = uint8(x * y)
			}
			pic[y] = row

		}
	}
	return pic

}

func main() {
	pic.Show(Pic)
}
