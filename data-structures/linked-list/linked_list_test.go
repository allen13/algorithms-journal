package linkedlist

import (
	"testing"
)

func TestLinkeList(t *testing.T) {
	ll := LinkedList{}

	ll.PushInt(0)
	ll.PushInt(1)

	val0 := ll.IntAt(0)
	if val0 != 0 {
		t.Errorf("expected %d, got %d", 0, val0)
	}

	val1 := ll.IntAt(1)
	if val1 != 1 {
		t.Errorf("expected %d, got %d", 1, val1)
	}
}
