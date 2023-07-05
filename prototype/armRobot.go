package prototype

import "fmt"


type ArmRobot struct {
	Arm string
	Children Robot
}


func (ar *ArmRobot)Clone() Robot {
	return &ArmRobot{
		Arm: ar.Arm,
		Children: ar.Children,
	}
}

func (ar *ArmRobot)Running() {
	    fmt.Printf("Robot Arm is %s",ar.Arm )
		if ar.Children != nil {
			ar.Children.Running()
		}
}