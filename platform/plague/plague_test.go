package plague

import "testing"

//para ejecutar testing, en terminal : go test ./platform/plague
//para ejecutar testing de todos: go test -cover ./...

func TestAdd(t *testing.T) {
	_plague := New()
	_plague.Add(Item{})
	if len(_plague.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	_plague := New()
	_plague.Add(Item{})
	result := _plague.GetAll()
	if len(result) != 1 {
		t.Errorf("Item was not added")
	}
}
