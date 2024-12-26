package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	trueBool1 := &Boolean{Value: true}
	trueBool2 := &Boolean{Value: true}
	falseBool1 := &Boolean{Value: false}
	falseBool2 := &Boolean{Value: false}

	if trueBool1.HashKey() != trueBool2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}

	if falseBool1.HashKey() != falseBool2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}

	if trueBool1.HashKey() == falseBool1.HashKey() {
		t.Errorf("booleans with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	int1 := &Integer{Value: 1}
	int2 := &Integer{Value: 1}
	int3 := &Integer{Value: 2}
	int4 := &Integer{Value: 2}

	if int1.HashKey() != int2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if int3.HashKey() != int4.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if int1.HashKey() == int3.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}
