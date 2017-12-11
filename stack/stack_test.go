package stack

import (
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	s := New()
	s.Push(420)
	if got := s.Len(); got != 1 {
		t.Errorf("Len() = %d, want 1", got)
	}
}

func TestPush(t *testing.T) {
	got := New()
	got.Push(420)
	want := &Stack{420}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Push() -> %v, want %v", got, want)
	}
}

func TestPop(t *testing.T) {
	want := 420
	s := &Stack{want}

	got, err := s.Pop()
	if err != nil {
		t.Errorf("Failed to pop %d from stack: %v", want, err)
	}

	if got != want {
		t.Errorf("Pop() = %d, want %d", got, want)
	}
}
