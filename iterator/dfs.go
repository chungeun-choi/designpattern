package iterator

import "fmt"

type DepthFirstSearch struct {
	//n int
	graph *Graph
	visited []bool
}
func (dfs *DepthFirstSearch)initalized(){
	
	for i :=0 ; i < len(dfs.graph.Arr); i++ {
		dfs.visited[i] = false
	}
}

func (dfs *DepthFirstSearch)GetNext(){
	dfs.dfs(1)
}

func (dfs *DepthFirstSearch) HashNext() bool {
	return true
}

func (dfs *DepthFirstSearch)dfs(v int) {
	dfs.visited[v] = true
	fmt.Print(v, " ")

	for i := 1; i <= len(dfs.visited); i++ {
		if dfs.graph.Arr[v][i] == 1 && !dfs.visited[i] {
			dfs.dfs(i)
		}
	}
}