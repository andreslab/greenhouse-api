package wet

import "testing"

//para ejecutar testing, en terminal : go test ./platform/wet
//para ejecutar testing de todos: go test -cover ./...

func TestAdd(t *testing.T) {
	_wet := New()
	_wet.Add(Item{})
	if len(_wet.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	_wet := New()
	_wet.Add(Item{})
	result := _wet.GetAll()
	if len(result) != 1 {
		t.Errorf("Item was not added")
	}
}
