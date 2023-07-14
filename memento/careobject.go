package memento


type CareObject struct {
	snapshot []*snapshot
}

func (co *CareObject) AddSnapShot(m *snapshot) {
	co.snapshot = append(co.snapshot, m)
}
 

func (co *CareObject) GetIndxSnapshot(index int) *snapshot{
	return co.snapshot[index]
}