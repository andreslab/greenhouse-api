package temperature

import "testing"

//para ejecutar testing, en terminal : go test ./platform/temperature
//para ejecutar testing de todos: go test -cover ./...

func TestAdd(t *testing.T) {
	_temperature := New()
	_temperature.Add(Item{})
	if len(_temperature.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	_temperature := New()
	_temperature.Add(Item{})
	result := _temperature.GetAll()
	if len(result) != 1 {
		t.Errorf("Item was not added")
	}
}
