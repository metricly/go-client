package core

type Element struct {
	fqn string
	name string
	samples []Sample
}

func (e *Element) AddSample(sample Sample) {
	e.samples = append(e.samples, sample)
}
