package checker

import "reflect"

const ResultRequired Result = "REQUIRED"

func checkRequired(value reflect.Value) Result {
	if value.IsZero() {
		return ResultRequired
	}

	k := value.Kind()

	if (k == reflect.Array || k == reflect.Map || k == reflect.Slice) && value.Len() == 0 {
		return ResultRequired
	}

	return ResultValid
}
