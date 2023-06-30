package observer

import "fmt"

type Store struct {
	subscriber []Observer
	product string
	isit bool
}

func NewStore(name string) *Store{
	return &Store{
		product: name,
	}
}

func (s *Store)UpdateProdcut(){
	fmt.Println("제품 입고됨!")
	s.isit = true
	s.NotifAll()
}

func (s *Store)Register(o Observer){
	s.subscriber = append(s.subscriber,o )
}

func (s *Store)Deregister(o Observer){
	s.subscriber = removeSubscriber(s.subscriber,o)
}

func (s *Store)NotifAll(){
	for _, value := range(s.subscriber){
		value.Update(s.product)
	}
		
}

func removeSubscriber(sub []Observer, outer Observer ) []Observer{
	observerLen := len(sub)

	for index,value := range(sub){
		if value.GetID() == outer.GetID(){
			sub[observerLen-1], sub[index] = sub[index], sub[observerLen-1]
            return sub[:observerLen-1]
		}
	}
	return sub
}
