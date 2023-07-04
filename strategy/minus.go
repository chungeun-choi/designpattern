package strategy


type Minus struct {
}


func(p *Minus) Operate(args []int) int {
	var result int
	
	for _,value := range(args){
		result -=value
	}

	return result
}