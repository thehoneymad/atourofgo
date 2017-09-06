package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := range ret {
		row := make([]uint8, dx)
		for j := range row {
			row[j] = uint8((i^j))
		}
		ret[i] = row
	}
	return ret
}

func main() {
	pic.Show(Pic)
}