package command

import "fmt"

type TvDevice struct {
	isRunning bool
}

func (t *TvDevice)on() {
	t.isRunning = true
	fmt.Printf("Turnning TV on")
}

func (t *TvDevice)off() {
	t.isRunning = false
	fmt.Printf("Turnning TV off")
}