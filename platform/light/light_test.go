package light

import "testing"

//para ejecutar testing, en terminal : go test ./platform/light
//para ejecutar testing de todos: go test -cover ./...

func TestAdd(t *testing.T) {
	_light := New()
	_light.Add(Item{})
	if len(_light.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	_light := New()
	_light.Add(Item{})
	result := _light.GetAll()
	if len(result) != 1 {
		t.Errorf("Item was not added")
	}
}
