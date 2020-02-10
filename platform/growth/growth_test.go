package growth

import "testing"

//para ejecutar testing, en terminal : go test ./platform/growth
//para ejecutar testing de todos: go test -cover ./...

func TestAdd(t *testing.T) {
	_growth := New()
	_growth.Add(Item{})
	if len(_growth.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	_growth := New()
	_growth.Add(Item{})
	result := _growth.GetAll()
	if len(result) != 1 {
		t.Errorf("Item was not added")
	}
}
