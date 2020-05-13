package raytrace

import (
	"fmt"
	"reflect"
	"testing"
)

func BresenhamSmokeTest() {
	for _, x_side := range []int{-1, 1} {
		for _, y_side := range []int{-1, 1} {
			resu := Bresenham2D(0, x_side*7, 0, y_side*4)

			for _, thingy := range resu {
				fmt.Println(thingy)
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
		t.Error(`Bresenham2d is different then expected!`)
	}

	//  t.Error(`IsPalindrome("detartrated") = false`)

}

/*
def test_bresenham2D():
  sx = 0
  sy = 1
  print("Testing bresenham2D...")
  r1 = bresenham2D(sx, sy, 10, 5)
  r1_ex = np.array([[0, 1, 2, 3, 4, 5, 6, 7, 8, 9,10],[1,1,2,2,3,3,3,4,4,5,5]])
  fig = plt.figure(figsize=(12,6))
  ax1 = fig.add_subplot(121)

  r2 = bresenham2D(sx, sy, 9, 6)
  r2_ex = np.array([[0,1,2,3,4,5,6,7,8,9],[1,2,2,3,3,4,4,5,5,6]])
  fig = plt.figure(figsize=(12,6))
  ax1 = fig.add_subplot(121)
  plt.scatter([sx,9],[sy,6],s=750,c='r',marker="s")
  plt.title("Given start and end point")
  plt.axis('equal')
  ax2 = fig.add_subplot(122)
  plt.scatter(r2_ex[0],r2_ex[1],s=750,c='b',marker="s")
  plt.title("bresenham2D return all the integer coordinates in-between")
  plt.axis('equal')
  plt.show()

  if np.logical_and(np.sum(r1 == r1_ex) == np.size(r1_ex),np.sum(r2 == r2_ex) == np.size(r2_ex)):
    print("...Test passed.")
  else:
    print("...Test failed.")

  # Timing for 1000 random rays
  num_rep = 1000
  start_time = time.time()
  for i in range(0,num_rep):
	  x,y = bresenham2D(sx, sy, 500, 200)
  print("1000 raytraces: --- %s seconds ---" % (time.time() - start_time))

*/
