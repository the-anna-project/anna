package round

// This file is generated by the CLG generator. Don't edit it manually. The CLG
// generator is invoked by go generate. For more information about the usage of
// the CLG generator check https://github.com/xh3b4sd/clggen or have a look at
// the clg package. There is the go generate statement placed to invoke clggen.

import (
	"reflect"
)

// filterError removes the last element of the given list. Thus filterError
// must only be used if the last element returned by a CLG implements the error
// interface. In case the last element is a non-nil error, this error is
// returned and the given list is discarded.
func filterError(values []reflect.Value) ([]reflect.Value, error) {
	if len(values) == 0 {
		return nil, nil
	}

	lastArg := values[len(values)-1]
	switch lastArg.Kind() {
	case reflect.Interface:
		fallthrough
	case reflect.Ptr:
		if err, ok := lastArg.Interface().(error); ok {
			return nil, maskAny(err)
		}
	}

	return values[:len(values)-1], nil
}
