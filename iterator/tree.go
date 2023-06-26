package iterator

type Tree struct {
	nodes []*Node 
	searchType string
}

func (t *Tree) CreateIterator() Iterator {
	if t.searchType == "bfs"{
		return &DepthFirstSearch{
			nodes: t.nodes,
		}
	}else if t.searchType == "dfs" {
		return &BreadthFirstSearch{
			nodes: t.nodes,
		}
	}else{
		panic("Not supported")
	}
}