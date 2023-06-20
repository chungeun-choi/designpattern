package builder

type Connectionbuilder interface {
	SetConfig()
	ConnetcObject()
	CheckConnection()
}


func GetbBuilder(name string) Connectionbuilder {
	if name == "elasticsearch" {
		return MakeEsBuilder()
	}else if name == "redis" {
		return MakeRedisBuilder()
	}else {
		panic("Not support to connection")
	}
}