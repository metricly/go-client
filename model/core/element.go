package core

type Element struct {
	Id string
	Name string
	Type string
	Location string
	attributes map[string]Attribute
	tags map[string]Tag
	relations map[string]Relation
	metrics map[string]Metric
	samples map[SampleKey]Sample
}

func NewElement(id string) Element {
	e := Element{}
	e.Id = id
	e.attributes = map[string]Attribute{}
	e.tags = map[string]Tag{}
	e.relations = map[string]Relation{}
	e.metrics = map[string]Metric{}
	e.samples = map[SampleKey]Sample{}
	return e
}

func (e *Element) AddAttribute(name, value string) {
	e.attributes[name] = Attribute{name, value}
}

func (e *Element) AddTag(name, value string) {
	e.tags[name] = Tag{name, value}
}

func (e *Element) AddRelation(fqn string) {
	e.relations[fqn] = Relation{fqn}
}

func (e *Element) AddMetric(metric Metric) {
	e.metrics[metric.Id] = metric
}

func (e *Element) AddSample(sample Sample) {
	e.samples[sample.Key()] = sample
	if _, found := e.metrics[sample.metricId]; !found {
		m := Metric{}
		m.Id = sample.metricId
		e.metrics[sample.metricId] = m
	}
}