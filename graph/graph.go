package graph

import "fmt"

type Edge struct {
	V int
	D int
}

//Adjacency lis
type Graph [][]int
type WGraph [][]Edge

func (g WGraph) String() string {
	var out string
	for i := range g {
		out += fmt.Sprintf("%d:\n\t", i)
		for j := 0; j < len(g[i]); j++ {
			if j != 0 {
				out += ", "
			}
			out += fmt.Sprintf("(V:%d, D:%d)", g[i][j].V, g[i][j].D)
		}
		out += "\n"
	}
	return out
}

func (g Graph) String() string {
	var out string
	for i := range g {
		out += fmt.Sprintf("%d:\n\t", i)
		out += fmt.Sprint( g[i])
		out += "\n"
	}
	return out
}

func SampleGraph() Graph {
	g := SampleGraphWeighted()
	out := make(Graph, len(g))
	for i := range g {
		out[i] = make([]int, len(g[i]))
		for j := range g[i] {
			out[i][j] = g[i][j].V
		}
	}
	return out
}

func SampleGraphWeighted() WGraph {
	return [][]Edge{
		{
			{1, 4},
			{7, 8},
		},
		{
			{7, 11},
			{2, 8},
		},
		{
			{3, 7},
			{7, 2},
			{1, 4},
		},
		{
			{2, 12},
			{5, 14},
			{4, 9},
		},
		{
			{0, 2},
			{5, 20},
			{5, 1},
		},
		{
			{6, 2},
			{7, 1},
		},
		{
			{1, 1},
			{3, 6},
		},
		{
			{7, 7},
		},
	}
}
