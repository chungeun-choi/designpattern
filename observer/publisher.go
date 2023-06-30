package observer

type Publisher interface {
    Register(observer Observer)
    Deregister(observer Observer)
    NotifAll()
}