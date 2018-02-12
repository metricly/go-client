package core

import "encoding/json"

type Metric struct {
	Id, Name, Unit, Type, SparseDataStrategy string
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

func (m *Metric) Tags() []Tag {
	tags := []Tag{}
	for _, t := range m.tags {
		tags = append(tags, t)
	}
	return tags
}

type MetricJSON struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
	Type string `json:"type"`
	SparseDataStrategy string `json:"sparseDataStrategy"`
	Tags []Tag `json:"tags"`
}

func (m Metric) MarshalJSON() ([]byte, error) {
	mjson := MetricJSON{
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
	var mjson MetricJSON
	if err := json.Unmarshal(b, &mjson); err != nil {
		return err
	}
	*m = NewMetric(mjson.Id, mjson.Name, mjson.Unit, mjson.Type, mjson.SparseDataStrategy)
	for _, tag := range mjson.Tags {
		m.AddTag(tag.Name, tag.Value)
	}
	return nil
}