package core

//A Check is used to determine the state or health of an infrastructure resource, service or application, e.g.
//	Check{"heartbeat", "serverA", 150}
type Check struct {
	Name string
	ElementId string
	TTL int
}
