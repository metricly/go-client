package core

import "encoding/json"

type Element struct {
	Id, Name, Type, Location string
	attributes map[string]Attribute
	tags map[string]Tag
	relations map[string]Relation
	metrics map[string]Metric
	samples map[SampleKey]Sample
}

func NewElement(Id, Name, Type, Location string) Element {
	e := Element{}
	e.Id = Id
	e.Name = Name
	e.Type = Type
	e.Location = Location
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

func (e *Element) Samples() []Sample {
	samples := []Sample{}
	for _, s := range e.samples {
		samples = append(samples, s)
	}
	return samples
}

type ElementJSON struct {
	Id, Name, Type, Location string
	Samples []Sample
}

func (e Element) MarshalJSON() ([]byte, error) {
	ejson := ElementJSON{}
	ejson.Id = e.Id
	ejson.Name = e.Name
	ejson.Type = e.Type
	ejson.Location = e.Location
	ejson.Samples = e.Samples()

	return json.Marshal(ejson)
}

func (e *Element) UnmarshalJSON(b []byte) error {
	var ejson ElementJSON
	if err := json.Unmarshal(b, &ejson); err != nil {
		return err
	}
	*e = NewElement(ejson.Id, ejson.Name, ejson.Type, ejson.Location)
	for _, sample := range ejson.Samples {
		e.AddSample(sample)
	}
	return nil
}