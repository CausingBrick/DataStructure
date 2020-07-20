package reconstruction

import "testing"

func TestReconstruction(t *testing.T) {
	pre := []int{1, 2, 3, 4, 5, 6, 7}
	in := []int{3, 2, 4, 1, 6, 5, 7}
	Reconstruction(pre, in)
}

/*
3
4
2
6
7
5
1
*/
