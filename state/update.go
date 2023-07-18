package state

import (
	"fmt"
)


type UpdateRequest struct {
	VendingMachine Machine
}


func (ur *UpdateRequest)Add() error {
	return fmt.Errorf("Unable to support actions you entered")
}

func (ur *UpdateRequest)Update() error {
	fmt.Println("'Update state' is running")
	ur.VendingMachine.SetState(ur.VendingMachine.updateItem)
	return nil	
}

func (ur *UpdateRequest)Delete() error {
	return fmt.Errorf("Unable to support actions you entered")
}