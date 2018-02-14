package core

import "encoding/json"

//Event represents an informational message that is associated with an Element at a point of time, use its Construction function NewEvent to create an Event
type Event struct {
	Source, Title, Type string
	timestamp           Timestamp
	tags                map[string]Tag
	data                ElementMessage
}

//ElementMessage represents an informational message on an Element
//	Level value is one of {"INFO", "WARNING", "CRITICAL"}
type ElementMessage struct {
	ElementId string `json:"elementId"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

//NewEvent constructs an Event given its Source, Title, Type, timestamp, ElementMessage and optional Tags
func NewEvent(source, title, etype string, timestamp interface{}, message ElementMessage, tags ...Tag) Event {
	e := Event{}
	ts, _ := TimestampValue(timestamp)
	e.Source, e.Title, e.Type, e.timestamp = source, title, etype, ts
	e.data = message
	e.tags = map[string]Tag{}
	for _, tag := range tags {
		e.tags[tag.Name] = tag
	}
	return e
}

//Add a Tag to Event
func (e *Event) AddTag(name, value string) *Event {
	e.tags[name] = Tag{name, value}
	return e
}

//Get all Tags of an Event
func (e *Event) Tags() []Tag {
	tags := []Tag{}
	for _, t := range e.tags {
		tags = append(tags, t)
	}
	return tags
}

type eventJSON struct {
	Source    string         `json:"source"`
	Title     string         `json:"title"`
	Type      string         `json:"type"`
	Timestamp int64          `json:"timestamp"`
	Tags      []Tag          `json:"tags"`
	Data      ElementMessage `json:"data"`
}

func (e Event) MarshalJSON() ([]byte, error) {
	ejson := eventJSON{
		e.Source,
		e.Title,
		e.Type,
		int64(e.timestamp),
		e.Tags(),
		e.data,
	}
	return json.Marshal(ejson)
}
