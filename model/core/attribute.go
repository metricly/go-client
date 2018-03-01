package core

//Attribute is a Name/Value pair that is generally used to describe the characteristics of an Element, e.g.
//	Attribute{"platform", "Linux"}, Attribute{"cpus", "8"}
type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
