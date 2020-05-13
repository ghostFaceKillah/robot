package main

import (
	"fmt"
	"robot/algo/raytrace"
)


func BresenhamSmokeTest(){
	for _, x_side := range []int{-1, 1} {
		for _, y_side := range []int{-1, 1} {
			resu := raytrace.Bresenham2D(0, x_side * 7, 0, y_side * 4)

			for _, thingy := range resu {
				fmt.Println(thingy)
			}

		}
	}
}


func BresenhamTranslationInvariance() {
	// here we will see if we transpose by some integer ammount
}

func main() {
	// ok kinda works




}

func printSlice(s []int) {
	fmt.Println("len=%d cap=%d %v\n", len(s), cap(s), s)
}
