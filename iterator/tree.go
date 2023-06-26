package iterator

type Tree struct {
	nodes *Graph 
	searchType string
}

func (t *Tree) CreateIterator() Iterator {
	if t.searchType == "bfs"{
		return &DepthFirstSearch{
			graph: t.nodes,
		}
	}else if t.searchType == "dfs" {
		return &BreadthFirstSearch{
			graph: t.nodes,
		}
	}else{
		panic("Not supported")
	}
}


