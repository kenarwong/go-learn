package drawpic

import (
	"golang.org/x/tour/pic"
)

func PicFunction(x, y uint8) uint8 {
	return (x + y) / 2
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)
	for x := range pic {
		row := make([]uint8, dy)
		for y := range row {
			row[y] = PicFunction(uint8(x), uint8(y))
			pic[x] = row
		}
	}

	return pic
}

func DrawPic() {
	pic.Show(Pic)
}
