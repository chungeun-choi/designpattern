package strategy


type Plus struct {
}


func(p *Plus) Operate(args []int) int {
	var result int
	
	for _,value := range(args){
		result +=value
	}

	return result
}