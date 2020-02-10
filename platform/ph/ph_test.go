package ph

import "testing"

//para ejecutar testing, en terminal : go test ./platform/ph
//para ejecutar testing de todos: go test -cover ./...

func TestAdd(t *testing.T) {
	_ph := New()
	_ph.Add(Item{})
	if len(_ph.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	_ph := New()
	_ph.Add(Item{})
	result := _ph.GetAll()
	if len(result) != 1 {
		t.Errorf("Item was not added")
	}
}
