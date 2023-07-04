package strategy


type Numbser struct {
	strategy CaculatorAlog
	Parameter []int
	
}

func (c * Numbser) SetStrategy(stObj CaculatorAlog){
	c.strategy = stObj

}


func (c* Numbser) DoSomething() int{
	return c.strategy.Operate(c.Parameter)
}