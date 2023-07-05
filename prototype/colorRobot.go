package prototype

import "fmt"

type ColorRobot struct {
	Color string
}


func (cr *ColorRobot)Clone() Robot {
	return &ColorRobot{
		Color: cr.Color,
	}
}


func (cr *ColorRobot)Running() {
	fmt.Printf("Robot Color is %s",cr.Color)
}

