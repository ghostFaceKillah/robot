package raytrace

// here's a good explaination of the Bresenham's algo
// https://www.youtube.com/watch?v=vlZFSzCIwoc
// some details are still kind of unclear to me, but OK

// what should be even a good list of outputs?
// it's kind of fine to return a list or sth

// it should really be built in ...
func Abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func Sign(x int) int {
	if x > 0 {
		return 1
	} else {
		return -1
	}
}

type Point struct {
	x int
	y int
}

func Bresenham2D(x_1, x_2, y_1, y_2 int) []Point {
	var resu []Point

	dx := Abs(x_2 - x_1)
	g := Sign(x_2 - x_1)

	dy := Abs(y_2 - y_1)
	h := Sign(y_2 - y_1)

	if dx > dy {
		c := -dx
		y := y_1
		for x := x_1; x != x_2; x += g {
			resu = append(resu, Point{x, y})
			c += 2 * dy
			if c > 0 {
				y += h
				c -= 2 * dx
			}
		}
	} else {
		c := -dy
		x := x_1
		for y := y_1; y != y_2; y += h {
			resu = append(resu, Point{x, y})
			c += 2 * dx
			if c > 0 {
				x += g
				c -= 2 * dy
			}
		}
	}

	return resu
}
