package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
)

// AMGraph Adjancency Matrix Graph.
type AMGraph struct {
	vexs           []int
	arcs           [][]int
	vexnum, arcnum int
}

func printMatrix(g *AMGraph) {
	fmt.Println("vexs:", g.vexs)
	fmt.Println("Matrix:")
	for _, v := range g.arcs {
		fmt.Println(v)
	}
}

// CreatUDN scan data to g
func CreatUDN(g *AMGraph) {
	// fmt.Println("Scaning vexnum arcnum:")
	fmt.Scanf("%d %d", &g.vexnum, &g.arcnum)
	initUDN(g, g.vexnum, g.arcnum)
	// fmt.Println("Scaning each vex:")
	for i := 0; i < g.vexnum; i++ {
		fmt.Scanf("%d", &g.vexs[i])
	}
	for i := 0; i < g.vexnum; i++ {
		for j := 0; j < g.vexnum; j++ {
			g.arcs[i][j] = -1
		}
	}
	// fmt.Println("Scaning each pair arcs and their weights:")
	for k := 0; k < g.arcnum; k++ {
		var v1, v2, w int
		fmt.Scanf("%d %d %d", &v1, &v2, &w)
		i, j := locateVex(g, v1), locateVex(g, v2)
		g.arcs[i][j], g.arcs[j][i] = w, w
	}
}
func initUDN(g *AMGraph, vn, an int) {
	g.vexnum, g.arcnum = vn, an
	g.vexs = make([]int, g.vexnum)
	g.arcs = make([][]int, g.vexnum)
	for i := 0; i < g.vexnum; i++ {
		g.arcs[i] = make([]int, g.vexnum)
	}
}
func locateVex(g *AMGraph, val int) int {
	for i := 0; i < g.vexnum; i++ {
		if g.vexs[i] == val {
			return i
		}
	}
	return -1
}

// DFS traverse graph by depth first from node v
func DFS(g *AMGraph, visited []bool, v int) {
	fmt.Println("->", v)
	i := locateVex(g, v)
	visited[i] = true
	for j := 0; j < g.vexnum; j++ {
		if g.arcs[i][j] != -1 && !visited[j] {
			DFS(g, visited, g.vexs[j])
		}
	}
}

// BFS traverse graph by breadth first from node v
func BFS(g *AMGraph, visited []bool, queue *list.List, v int) {
	fmt.Println("->", v)
	visited[locateVex(g, v)] = true
	queue.PushFront(v)
	for queue.Len() != 0 {
		u := queue.Remove(queue.Front()).(int)
		for i := 0; i < g.vexnum; i++ {
			if g.arcs[locateVex(g, u)][i] != -1 && !visited[i] {
				fmt.Println("->", g.vexs[i])
				visited[i] = true
				queue.PushFront(g.vexs[i])
			}
		}
	}
}

type element struct {
	adjvex  int
	lowcost int
}

// Prim create a minimum cost spanning tree.
func Prim(g *AMGraph, closedge []element, v int) {
	k := locateVex(g, v)
	for i := 0; i < g.vexnum; i++ {
		if i != k {
			closedge[i] = element{v, g.arcs[k][i]}
		}
	}
	closedge[k].lowcost = 0
	for i := 1; i < g.vexnum; i++ {
		k = minimumCost(closedge)
		// println(k)
		u0, v0 := closedge[k].adjvex, g.vexs[k]
		// fmt.Println(closedge)
		fmt.Printf("%d<->%d\n", u0, v0)
		closedge[k].lowcost = 0
		for j := 0; j < g.vexnum; j++ {
			if g.arcs[k][j] == -1 {
				continue
			}
			if g.arcs[k][j] < closedge[j].lowcost || closedge[j].lowcost == -1 {
				// fmt.Println(g.arcs[k][j], closedge[j].lowcost)
				closedge[j] = element{g.vexs[k], g.arcs[k][j]}
			}
		}
	}
}

func minimumCost(closedge []element) int {
	index := -1
	mincost := math.Inf(0)
	for i, v := range closedge {
		if v.lowcost <= 0 {
			continue
		}
		if mincost > float64(v.lowcost) {
			index, mincost = i, float64(v.lowcost)
		}
	}
	return index
}

// Edge stores two vexes and their weight.
type Edge []struct {
	head, tail int
	lowcost    int
}

func (e Edge) Len() int {
	return len(e)
}
func (e Edge) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e Edge) Less(i, j int) bool {
	return e[i].lowcost < e[j].lowcost
}

// Kruskal create a minimum cost spanning tree.
func Kruskal(g *AMGraph, edge Edge, vexset []int) {
	sort.Sort(edge)
	for i := 0; i < g.vexnum; i++ {
		vexset[i] = i
	}
	for i := 0; i < g.arcnum; i++ {
		v1, v2 := edge[i].head, edge[i].tail
		vs1, vs2 := vexset[v1], vexset[v2]
		if vs1 != vs2 {
			fmt.Printf("%d<->%d\n", g.vexs[edge[i].head], g.vexs[edge[i].tail])
			for j := 0; j < g.vexnum; j++ {
				if vexset[j] == vs2 {
					vexset[j] = vs1
				}
			}
		}
	}
}

func main() {
	var g1 AMGraph
	CreatUDN(&g1)
	printMatrix(&g1)
	fmt.Println("DFS:")
	visited := make([]bool, g1.vexnum)
	DFS(&g1, visited, 4)

	fmt.Println("BFS:")
	visited = make([]bool, g1.vexnum)
	var queue = list.New()
	BFS(&g1, visited, queue, 4)

	fmt.Println("Prim:")
	closedge := make([]element, g1.vexnum)
	Prim(&g1, closedge, 1)

	fmt.Println("Kruskal:")
	var edge Edge
	for i := 0; i < g1.vexnum; i++ {
		for j := i; j < g1.vexnum; j++ {
			if g1.arcs[i][j] != -1 {
				edge = append(edge, struct {
					head    int
					tail    int
					lowcost int
				}{i, j, g1.arcs[i][j]})
			}
		}
	}
	sort.Sort(edge)
	vexset := make([]int, g1.vexnum)
	Kruskal(&g1, edge, vexset)

}

/*
$ go build && cat test_data2| ./AdjacencyMatrix
!-output:-!
vexs: [1 2 3 4 5 6]
Matrix:
[-1 6 1 5 -1 -1]
[6 -1 5 -1 3 -1]
[1 5 -1 5 6 4]
[5 -1 5 -1 -1 2]
[-1 3 6 -1 -1 6]
[-1 -1 4 2 6 -1]
DFS:
-> 4
-> 1
-> 2
-> 3
-> 5
-> 6
BFS:
-> 4
-> 1
-> 3
-> 6
-> 5
-> 2
Prim:
1<->3
3<->6
6<->4
3<->2
2<->5
Kruskal:
1<->3
4<->6
2<->5
3<->6
2<->3
*/
