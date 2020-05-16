package raytrace

import (
	"fmt"
	"reflect"
	"testing"
)

// should make sure Bresenham is anti-symetric
func TestBresenhamSmoke(t *testing.T) {
	for _, x_side := range []int{-1, 1} {
		for _, y_side := range []int{-1, 1} {
			resu := Bresenham2D(0, x_side*7, 0, y_side*4)

			for _, thingy := range resu {
				fmt.Println(thingy)
			}

		}
	}
}

func showDifference(expected, real []Point) {
	if len(expected) != len(real) {
		fmt.Printf("Lengths are different: real = %d, expected = %d", len(real), len(expected))
	} else {
		// print the difference in elements
		for i := 0; i < len(real); i++ {
			e := expected[i]
			r := real[i]
			fmt.Printf("Expected= {X:%d, Y:%d}, got = {X:%d Y:%d}\n", e.X, e.Y, r.X, r.Y)
			if reflect.DeepEqual(e, r) {
				fmt.Println("These are equal")
			} else {
				fmt.Println("These are not equal")
			}
		}
	}
}

func TestBresenham2DvsExpected(t *testing.T) {
	expected := []Point{
		Point{0, 1}, Point{1, 1}, Point{2, 2}, Point{3, 2},
		Point{4, 3}, Point{5, 3}, Point{6, 3}, Point{7, 4},
		Point{8, 4}, Point{9, 5}, Point{10, 5},
	}

	real := Bresenham2D(0, 10, 1, 5)

	if !reflect.DeepEqual(expected, real) {
		showDifference(expected, real)
		t.Error(`Bresenham2d is different then expected!`)
	}
}

func BresenhamTranslationInvariance() {
	// here we will see if we transpose by some integer ammount
}
