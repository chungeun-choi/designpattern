package state


type StateInterface interface {
	Add() error
	Update() error
	Delete() error
}