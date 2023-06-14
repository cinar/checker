package checker

import (
	"reflect"
	"testing"
)

func TestCheckRequiredValidString(t *testing.T) {
	s := "valid"

	if checkRequired(reflect.ValueOf(s), reflect.ValueOf(nil)) != ResultValid {
		t.Fail()
	}
}

func TestCheckRequiredUninitializedString(t *testing.T) {
	var s string

	if checkRequired(reflect.ValueOf(s), reflect.ValueOf(nil)) != ResultRequired {
		t.Fail()
	}
}

func TestCheckRequiredEmptyString(t *testing.T) {
	s := ""

	if checkRequired(reflect.ValueOf(s), reflect.ValueOf(nil)) != ResultRequired {
		t.Fail()
	}
}

func TestCheckRequiredUninitializedNumber(t *testing.T) {
	var n int

	if checkRequired(reflect.ValueOf(n), reflect.ValueOf(nil)) != ResultRequired {
		t.Fail()
	}
}

func TestCheckRequiredValidSlice(t *testing.T) {
	s := []int{1}

	if checkRequired(reflect.ValueOf(s), reflect.ValueOf(nil)) != ResultValid {
		t.Fail()
	}
}

func TestCheckRequiredUninitializedSlice(t *testing.T) {
	var s []int

	if checkRequired(reflect.ValueOf(s), reflect.ValueOf(nil)) != ResultRequired {
		t.Fail()
	}
}

func TestCheckRequiredEmptySlice(t *testing.T) {
	s := make([]int, 0)

	if checkRequired(reflect.ValueOf(s), reflect.ValueOf(nil)) != ResultRequired {
		t.Fail()
	}
}

func TestCheckRequiredValidArray(t *testing.T) {
	s := [1]int{1}

	if checkRequired(reflect.ValueOf(s), reflect.ValueOf(nil)) != ResultValid {
		t.Fail()
	}
}

func TestCheckRequiredEmptyArray(t *testing.T) {
	s := [1]int{}

	if checkRequired(reflect.ValueOf(s), reflect.ValueOf(nil)) != ResultRequired {
		t.Fail()
	}
}

func TestCheckRequiredValidMap(t *testing.T) {
	m := map[string]string{
		"a": "b",
	}

	if checkRequired(reflect.ValueOf(m), reflect.ValueOf(nil)) != ResultValid {
		t.Fail()
	}
}

func TestCheckRequiredUninitializedMap(t *testing.T) {
	var m map[string]string

	if checkRequired(reflect.ValueOf(m), reflect.ValueOf(nil)) != ResultRequired {
		t.Fail()
	}
}

func TestCheckRequiredEmptyMap(t *testing.T) {
	m := map[string]string{}

	if checkRequired(reflect.ValueOf(m), reflect.ValueOf(nil)) != ResultRequired {
		t.Fail()
	}
}
