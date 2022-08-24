package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valResult := val.Call(nil)
		for _, res := range valResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	// return the value of a given variable
	val := reflect.ValueOf(x)

	// can't use NumField on pointers so the value has to be extracted using Elem()
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
