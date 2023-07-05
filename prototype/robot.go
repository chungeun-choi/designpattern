package prototype

type Robot interface {
	Clone() Robot
	Running()
}