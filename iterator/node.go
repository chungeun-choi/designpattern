package iterator

type Graph struct {
	Arr [4][4]int
}

// func (n *Graph) MakeNode(length int) {
// 	n.Arr = make([][]int, length)

// 	for i := range n.Arr {
// 		n.Arr[i] = make([]int, length)
// 	}
// }

func (n *Graph) InsertValueInNode(){
	//n.MakeNode(4)
	n.Arr = [4][4]int{{1,0,0,1},{0,1,1,0},{0,1,1,0},{1,0,0,1}}
	/*
	1001
	0110
	0110
	1001

	[[1,0,0,1],[0,1,1,0],[0,1,1,0],[1,0,0,1]]
	*/
	
}