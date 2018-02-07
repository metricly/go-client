package core

type Metric struct {
	Id string
	Name string
	Unit string
	Type string
	SparseDataStrategy string
	tags map[string]Tag
}

func NewMetric(Id, Name, Unit, Type, SparseDataStrategy string, tags ...Tag) Metric {
	m := Metric{}
	m.Id, m.Name, m.Unit, m.Type, m.SparseDataStrategy = Id, Name, Unit, Type, SparseDataStrategy
	m.tags = map[string]Tag{}
	for _, tag := range tags {
		m.tags[tag.Name] = tag
	}
	return m
}

func (m *Metric) AddTag(name, value string) {
	m.tags[name] = Tag{name, value}
}
