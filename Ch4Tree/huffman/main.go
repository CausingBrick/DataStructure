package main

import (
	"fmt"
)

type htNode struct {
	weight                 int
	parent, lchild, rchild int
}

func scanHuffmanValue() []htNode {
	var length int
	fmt.Scanf("%d", &length)
	ht := make([]htNode, 2*length)
	for i := 1; i < length+1; i++ {
		fmt.Scanf("%d", &ht[i].weight)
	}
	return ht
}
func smallest(ht []htNode) func(int) (int, int) {
	selectetd := make([]bool, len(ht))
	return func(n int) (int, int) {
		getmin := func() int {
			var min int
			getfirst := true
			for i := 1; i < n; i++ {
				if selectetd[i] {
					continue
				}
				if getfirst {
					min = i
					getfirst = false
					continue
				}
				if ht[min].weight > ht[i].weight {
					min = i
				}
			}
			selectetd[min] = true
			return min
		}
		return getmin(), getmin()
	}
}
func creatHuffmantree(ht []htNode) {
	sm := smallest(ht)
	for i := len(ht)/2 + 1; i < len(ht); i++ {
		s1, s2 := sm(i)
		ht[s1].parent = i
		ht[s2].parent = i
		ht[i].lchild = s1
		ht[i].rchild = s2
		ht[i].weight = ht[s1].weight + ht[s2].weight
	}
}

func main() {
	ht := scanHuffmanValue()
	creatHuffmantree(ht)
	for i := 1; i < len(ht); i++ {
		fmt.Println(i, ht[i])
	}
}
