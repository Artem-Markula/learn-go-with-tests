package main

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	Tone := new(SingleTone)
	
	object1 := Tone.GetInstance()
	object2 := Tone.GetInstance()

	if object1 != object2 {
		t.Error("This not equal\n")
	}
}
