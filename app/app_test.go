package app

import "testing"

func TestApp_Initialize(t *testing.T) {
	a := App{}
	a.Initialize()

	if a.DB == nil {
		t.Errorf("Database wasn't instanced.")
	} else if a.Router == nil {
		t.Errorf("Routes weren't declared.")
	}
}

