package state

import (
	"fmt"
)


type DeleteRequest struct {
	VendingMachine Machine
}

func (dr *DeleteRequest)Add() error {
	return fmt.Errorf("Unable to support actions you entered")
}

func (dr *DeleteRequest)Update() error {
	return fmt.Errorf("Unable to support actions you entered")
}

func (dr *DeleteRequest)Delete() error {
	fmt.Println("'Delete state' is running")
	dr.VendingMachine.SetState(dr.VendingMachine.deleteItem)
	return nil
}