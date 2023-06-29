package observer

type Publisher interface {
    register(observer Observer)
    deregister(observer Observer)
    notifyAll()
}