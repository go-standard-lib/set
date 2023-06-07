package set

import "testing"

type testStruct struct {
	test bool
}

func TestSetNew(t *testing.T) {
	set := New[string]()

	if set == nil {
		t.Fatal("Failed to create set")
	}
}

func TestSetAdd(t *testing.T) {
	set := New[string]()

	value := "hello"

	set.Add(&value)

	if !set.Contains(&value) {
		t.Fatal("Value", value, "was not added to the set")
	}
}

func TestRemoveAdd(t *testing.T) {
	set := New[string]()

	value := "hello"

	set.Add(&value)

	set.Remove(&value)

	if set.Contains(&value) {
		t.Fatal("Value", value, "was not removed to the set")
	}
}

func TestAddStruct(t *testing.T) {
	set := New[testStruct]()

	value := &testStruct{
		test: true,
	}

	value2 := &testStruct{
		test: false,
	}

	value3 := value

	set.Add(value)
	set.Add(value2)
	set.Add(value3)

	if !set.Contains(value) {
		t.Fatal("Value", value, "was not added to the set")
	}

	if !set.Contains(value2) {
		t.Fatal("Value", value2, "was not added to the set")
	}

	if set.Size() != 2 {
		t.Fatal("A value was added as a copy instead of a pointer")
	}
}
