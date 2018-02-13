package core

//Tag is a Name/Value pair that is associated with an Element, Metric or Event, generally for grouping purpose, e.g.
//	Tag{"Group", "uat"}
type Tag struct {
	Name string `json:"name"`
	Value string `json:"value"`
}
