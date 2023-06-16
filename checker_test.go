package checker

import (
	"reflect"
	"testing"
)

func TestInitCheckersUnknown(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	initCheckers("unknown")
}

func TestInitCheckersKnwon(t *testing.T) {
	checkers := initCheckers("required")

	if len(checkers) != 1 {
		t.Fail()
	}

	if reflect.ValueOf(checkers[0]).Pointer() != reflect.ValueOf(checkRequired).Pointer() {
		t.Fail()
	}
}

func TestRegister(t *testing.T) {
	var checker CheckFunc = func(_, _ reflect.Value) Result {
		return ResultValid
	}

	var maker MakeFunc = func(_ string) CheckFunc {
		return checker
	}

	name := "test"

	Register(name, maker)

	checkers := initCheckers(name)

	if len(checkers) != 1 {
		t.Fail()
	}

	if reflect.ValueOf(checkers[0]).Pointer() != reflect.ValueOf(checker).Pointer() {
		t.Fail()
	}
}

func TestCheckInvalid(t *testing.T) {
	type Person struct {
		Name string `checkers:"required"`
	}

	person := &Person{}

	mistakes, valid := Check(person)
	if valid {
		t.Fail()
	}

	if len(mistakes) != 1 {
		t.Fail()
	}

	if mistakes["Name"] != ResultRequired {
		t.Fail()
	}
}

func TestCheckValid(t *testing.T) {
	type Person struct {
		Name string `checkers:"required"`
	}

	person := &Person{
		Name: "Onur",
	}

	mistakes, valid := Check(person)
	if !valid {
		t.Fail()
	}

	if len(mistakes) != 0 {
		t.Fail()
	}
}

func TestCheckNoStruct(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	s := "unknown"
	Check(s)
}

func TestCheckNestedStruct(t *testing.T) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name string `checkers:"required"`
		Home Address
	}

	person := &Person{}

	mistakes, valid := Check(person)
	if valid {
		t.Fail()
	}

	if len(mistakes) != 2 {
		t.Fail()
	}

	if mistakes["Name"] != ResultRequired {
		t.Fail()
	}

	if mistakes["Home.Street"] != ResultRequired {
		t.Fail()
	}
}

func TestNumberOfInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	s := "invalid"

	numberOf(reflect.ValueOf(s))
}

func TestNumberOfInt(t *testing.T) {
	n := 10

	if numberOf(reflect.ValueOf(n)) != float64(n) {
		t.Fail()
	}
}

func TestNumberOfFloat(t *testing.T) {
	n := float64(10.10)

	if numberOf(reflect.ValueOf(n)) != n {
		t.Fail()
	}
}

func BenchmarkCheck(b *testing.B) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name string `checkers:"required"`
		Home Address
	}

	person := &Person{}

	for n := 0; n < b.N; n++ {
		Check(person)
	}
}
