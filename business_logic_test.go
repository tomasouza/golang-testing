package golang_testing

import "testing"

func TestBusinessLogic(t *testing.T) {
	ht := NewInMemoryHashTable()
	BusinessLogic(ht)
	val, err := ht.Get("hello")
	if err != nil {
		t.Fatalf("error on Get: %s", err)
	}
	if string(val) != "world" {
		t.Fatalf("expected 'world', got %s", val)
	}
}
