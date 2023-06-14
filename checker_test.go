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
