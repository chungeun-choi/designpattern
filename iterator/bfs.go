package iterator

import "fmt"

type BreadthFirstSearch struct {
	n int 
	graph *Graph
	visited []bool
	queue []int
}

func (bfs *BreadthFirstSearch)initalized(){
	for i :=0 ; i < len(bfs.graph.Arr); i++ {
		bfs.visited[i] = false
	}
}

func (bfs *BreadthFirstSearch)GetNext(){
	bfs.bfs(1)
}

func (bfs *BreadthFirstSearch)HashNext() bool{
	return true
}

func (bfs *BreadthFirstSearch)bfs(v int) {
	bfs.visited[v] = true
	queue := []int{v}

	for len(queue) != 0 {
		front := queue[0]
		fmt.Print(front, " ")
		queue = queue[1:]

		for i := 1; i <= len(bfs.visited) ; i++ {
			if bfs.graph.Arr[front][i] == 1 && !bfs.visited[i] {
				bfs.visited[i] = true
				queue = append(queue, i)
			}
		}
	}
}