package greet

import "testing"

func TestMorning(t *testing.T) {
	if Morning != "Hey, Good morning" {
		t.Errorf("Morning == %q", "Hey, Good morning")
	}
}