package core

import "encoding/json"

//Element is a physical entity such as a server or a logical component such as a transaction, use its Construction function NewElement to create an Element
type Element struct {
	Id, Name, Type, Location string
	attributes               map[string]Attribute
	tags                     map[string]Tag
	relations                map[string]Relation
	metrics                  map[string]Metric
	samples                  map[sampleKey]Sample
}

//NewElement constructs an Element given its Id, Name, Type and Location
func NewElement(id, name, etype, location string) Element {
	e := Element{}
	e.Id = id
	e.Name = name
	e.Type = etype
	e.Location = location
	e.attributes = map[string]Attribute{}
	e.tags = map[string]Tag{}
	e.relations = map[string]Relation{}
	e.metrics = map[string]Metric{}
	e.samples = map[sampleKey]Sample{}
	return e
}

//Add an Attribute to Element
func (e *Element) AddAttribute(name, value string) *Element {
	e.attributes[name] = Attribute{name, value}
	return e
}

//Get all Attributes of an Element
func (e *Element) Attributes() []Attribute {
	attributes := []Attribute{}
	for _, a := range e.attributes {
		attributes = append(attributes, a)
	}
	return attributes
}

func (e *Element) Attribute(name string) (Attribute, bool) {
	attribute, ok := e.attributes[name]
	return attribute, ok
}

//Add a Tag to Element
func (e *Element) AddTag(name, value string) *Element {
	e.tags[name] = Tag{name, value}
	return e
}

//Get all Tags of an Element
func (e *Element) Tags() []Tag {
	tags := []Tag{}
	for _, t := range e.tags {
		tags = append(tags, t)
	}
	return tags
}

//Get a Tag by its name
func (e *Element) Tag(name string) (Tag, bool) {
	tag, ok := e.tags[name]
	return tag, ok
}

//Add a Relation to Element
func (e *Element) AddRelation(fqn string) *Element {
	e.relations[fqn] = Relation{fqn}
	return e
}

//Get all Relations of an Element
func (e *Element) Relations() []Relation {
	relations := []Relation{}
	for _, r := range e.relations {
		relations = append(relations, r)
	}
	return relations
}

//Add a Metric to Element
func (e *Element) AddMetric(metric Metric) *Element {
	e.metrics[metric.Id] = metric
	return e
}

//Get all Metrics of an Element
func (e *Element) Metrics() []Metric {
	metrics := []Metric{}
	for _, m := range e.metrics {
		metrics = append(metrics, m)
	}
	return metrics
}

//Add a Sample to Element
func (e *Element) AddSample(sample Sample) *Element {
	e.samples[sample.key()] = sample
	if _, found := e.metrics[sample.metricId]; !found {
		m := Metric{}
		m.Id = sample.metricId
		m.SparseDataStrategy = "None"
		e.metrics[sample.metricId] = m
	}
	return e
}

//Get all Samples of an Element
func (e *Element) Samples() []Sample {
	samples := []Sample{}
	for _, s := range e.samples {
		samples = append(samples, s)
	}
	return samples
}

type elementJSON struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Location   string      `json:"location"`
	Attributes []Attribute `json:"attributes"`
	Tags       []Tag       `json:"tags"`
	Relations  []Relation  `json:"relations"`
	Metrics    []Metric    `json:"metrics"`
	Samples    []Sample    `json:"samples"`
}

func (e Element) MarshalJSON() ([]byte, error) {
	ejson := elementJSON{
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
	var ejson elementJSON
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
