package state

import "fmt"

type AddRequest struct {
	VendingMachine Machine
}

func (ar *AddRequest)Add() error{
	fmt.Println("'Add state' is runnning ")
	ar.VendingMachine.SetState(ar.VendingMachine.addItem)
	return nil
}

func (ar *AddRequest)Update() error {
	return fmt.Errorf("Unable to support actions you entered")
}

func (ar *AddRequest)Delete() error{
	return fmt.Errorf("Unable to support actions you entered") 
}