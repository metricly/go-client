package core

import "testing"

func TestMetric(t *testing.T) {
	//given
	test := struct {
		Id, Name, Unit, Type, SparseDataStrategy string
	} {
		Id: "cpu.percent",
		Name: "CPU Percent",
		Unit: "%",
		Type: "COUNTER",
		SparseDataStrategy: "NONE",
	}
	//when
	m := NewMetric(test.Id, test.Name, test.Unit, test.Type, test.SparseDataStrategy, Tag{"env", "prod"})
	//then
	if m.Id != test.Id {
		t.Errorf("metric Id expected: %s, actual: %s", test.Id, m.Id)
	}
	if m.Name != test.Name {
		t.Errorf("metric Name expected: %s, actual: %s", test.Name, m.Name)
	}
	if m.Unit != test.Unit {
		t.Errorf("metric Unit expected: %s, actual: %s", test.Unit, m.Unit)
	}
	if m.Type != test.Type {
		t.Errorf("metric Type expected: %s, actual: %s", test.Type, m.Type)
	}
	if m.SparseDataStrategy != test.SparseDataStrategy {
		t.Errorf("metric SparseDataStrategy expected: %s, actual: %s", test.SparseDataStrategy, m.SparseDataStrategy)
	}
	if len(m.tags) != 1 {
		t.Errorf("metric should contain exactly 1 tag")
	}
	if _, found := m.tags["env"]; !found {
		t.Errorf("metric should contain a tag with key 'env'")
	}
	if tag := m.tags["env"]; tag.Value != "prod" {
		t.Errorf("metric's tag 'env' should have value 'prod'")
	}
}
