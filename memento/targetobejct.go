package memento

type TargetObject struct {
	Field1 string
	Field2 string
	Field3 string 
}


func (to *TargetObject) CreateSanpshot() *snapshot{
	return &snapshot{
		field1: to.Field1,
		field2: to.Field2,
		field3: to.Field3,
	}
}	

func (to *TargetObject)ResotreSnapshot(s *snapshot) {
	

	to.Field1 = s.field1
	to.Field2 = s.field2
	to.Field3 = s.field3
}
