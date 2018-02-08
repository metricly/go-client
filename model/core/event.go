package core

import "encoding/json"

type Event struct {
	Source, Title, Type string
	timestamp Timestamp
	tags map[string]Tag
	data ElementMessage
}

type ElementMessage struct {
	ElementId string
	Level string `json:"level"`
	Message string `json:"message"`
}

func NewEvent(Source, Title, Type string, timestamp interface{}, message ElementMessage, tags ...Tag) Event {
	e := Event{}
	ts, _ := TimestampValue(timestamp)
	e.Source, e.Title, e.Type, e.timestamp = Source, Title, Type, ts
	e.data = message
	e.tags = map[string]Tag{}
	for _, tag := range tags {
		e.tags[tag.Name] = tag
	}
	return e
}

func (e *Event) AddTag(name, value string) {
	e.tags[name] = Tag{name, value}
}

func (e *Event) Tags() []Tag {
	tags := []Tag{}
	for _, t := range e.tags {
		tags = append(tags, t)
	}
	return tags
}

type EventJSON struct {
	Source string `json:"source"`
	Title string `json:"title"`
	Type string `json:"type"`
	Timestamp int64 `json:"timestamp"`
	Tags []Tag `json:"tags"`
	Data ElementMessage `json:"data"`
}

func (e Event) MarshalJSON() ([]byte, error) {
	ejson := EventJSON{
		e.Source,
		e.Title,
		e.Type,
		int64(e.timestamp),
		e.Tags(),
		e.data,
	}
	return json.Marshal(ejson)
}