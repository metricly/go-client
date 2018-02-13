package core

import "encoding/json"

//Metric is a quantifiable measurement that is associated with an Element, use its Construction function NewMetric(...) to create a Metric
//	SparseDataStrategy value is one of {"None", "ReplaceWithZero", "ReplaceWithLast", "ReplaceWithHistoricalMax", "ReplaceWithHistoricalMin"}
type Metric struct {
	Id, Name, Unit, Type, SparseDataStrategy string
	tags map[string]Tag
}

//NewMetric constructs a Metric given its Id, Name, Unit, Type, SparseDataStrategy and optional Tags
func NewMetric(id, name, unit, mtype, sparseDataStrategy string, tags ...Tag) Metric {
	m := Metric{}
	m.Id, m.Name, m.Unit, m.Type, m.SparseDataStrategy = id, name, unit, mtype, sparseDataStrategy
	m.tags = map[string]Tag{}
	for _, tag := range tags {
		m.tags[tag.Name] = tag
	}
	return m
}

//Add a Tag to Metric
func (m *Metric) AddTag(name, value string) *Metric {
	m.tags[name] = Tag{name, value}
	return m
}

//Get all Tags of a Metric
func (m *Metric) Tags() []Tag {
	tags := []Tag{}
	for _, t := range m.tags {
		tags = append(tags, t)
	}
	return tags
}

type metricJSON struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
	Type string `json:"type"`
	SparseDataStrategy string `json:"sparseDataStrategy"`
	Tags []Tag `json:"tags"`
}

func (m Metric) MarshalJSON() ([]byte, error) {
	mjson := metricJSON{
		m.Id,
		m.Name,
		m.Unit,
		m.Type,
		m.SparseDataStrategy,
		m.Tags(),
	}
	return json.Marshal(mjson)
}

func (m *Metric)  UnmarshalJSON(b []byte) error {
	var mjson metricJSON
	if err := json.Unmarshal(b, &mjson); err != nil {
		return err
	}
	*m = NewMetric(mjson.Id, mjson.Name, mjson.Unit, mjson.Type, mjson.SparseDataStrategy)
	for _, tag := range mjson.Tags {
		m.AddTag(tag.Name, tag.Value)
	}
	return nil
}