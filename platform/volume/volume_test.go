package volume

import "testing"

//para ejecutar testing, en terminal : go test ./platform/volume
//para ejecutar testing de todos: go test -cover ./...

func TestAdd(t *testing.T) {
	_volume := New()
	_volume.Add(Item{})
	if len(_volume.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	_volume := New()
	_volume.Add(Item{})
	result := _volume.GetAll()
	if len(result) != 1 {
		t.Errorf("Item was not added")
	}
}
