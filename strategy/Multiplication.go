package strategy


type Multiplication struct {
	
}


func (m *Multiplication) Operate(args []int) int{
	var result int
	for _,value := range(args){
		result *=value
	}

	return result
}