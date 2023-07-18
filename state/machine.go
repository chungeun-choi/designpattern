package state

type Machine struct {
    addItem       StateInterface
    deleteItem StateInterface
    updateItem      StateInterface

    currentState StateInterface

}


func NewMachine() *Machine{
	machine := &Machine{

	}

	addRequest := &AddRequest{
		VendingMachine: *machine,
	}

	deleteRequest := &DeleteRequest{
		VendingMachine: *machine,
	}

	updateRequest := &UpdateRequest{
		VendingMachine: *machine,
	}


	machine.currentState = addRequest
	machine.addItem = addRequest
	machine.deleteItem = updateRequest
	machine.updateItem = deleteRequest


	return machine
	
}


func (m *Machine)Add() error {
	return m.addItem.Add()
}

func  (m *Machine)Update()error {
	return m.updateItem.Update()
}

func  (m *Machine)Delete()error{
	return m.deleteItem.Delete()
}

func (m *Machine)SetState(state StateInterface){
	 m.currentState = state
}