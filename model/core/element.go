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

func (e *Element) Attributes() []Attribute {
	attributes := []Attribute{}
	for _, a := range e.attributes {
		attributes = append(attributes, a)
	}
	return attributes
}

func (e *Element) AddTag(name, value string) {
	e.tags[name] = Tag{name, value}
}

func (e *Element) Tags() []Tag {
	tags := []Tag{}
	for _, t := range e.tags {
		tags = append(tags, t)
	}
	return tags
}

func (e *Element) AddRelation(fqn string) {
	e.relations[fqn] = Relation{fqn}
}

func (e *Element) Relations() []Relation {
	relations := []Relation{}
	for _, r :=  range e.relations {
		relations = append(relations, r)
	}
	return relations
}

func (e *Element) AddMetric(metric Metric) {
	e.metrics[metric.Id] = metric
}

func (e *Element) Metrics() []Metric {
	metrics := []Metric{}
	for _, m := range e.metrics {
		metrics = append(metrics, m)
	}
	return metrics
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
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Location string `json:"location"`
	Attributes []Attribute `json:"attributes"`
	Tags []Tag `json:"tags"`
	Relations []Relation `json:"relations"`
	Metrics []Metric `json:"metrics"`
	Samples []Sample `json:"samples"`
}

func (e Element) MarshalJSON() ([]byte, error) {
	ejson := ElementJSON{
		e.Id,
		e.Name,
		e.Type,
		e.Location,
		e.Attributes(),
		e.Tags(),
		e.Relations(),
		e.Metrics(),
		e.Samples(),
	}

	return json.Marshal(ejson)
}

func (e *Element) UnmarshalJSON(b []byte) error {
	var ejson ElementJSON
	if err := json.Unmarshal(b, &ejson); err != nil {
		return err
	}
	*e = NewElement(ejson.Id, ejson.Name, ejson.Type, ejson.Location)
	for _, attribute := range ejson.Attributes {
		e.AddAttribute(attribute.Name, attribute.Value)
	}
	for _, tag := range ejson.Tags {
		e.AddTag(tag.Name, tag.Value)
	}
	for _, relation := range ejson.Relations {
		e.AddRelation(relation.Fqn)
	}
	for _, metric := range ejson.Metrics {
		e.AddMetric(metric)
	}
	for _, sample := range ejson.Samples {
		e.AddSample(sample)
	}
	return nil
}