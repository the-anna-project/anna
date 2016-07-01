package findconnections

import (
	"reflect"
)

// filterError expects the given list of relfect values to be the output of a
// CLG execution. In case a CLG returns an error, the error must be the last
// element of the output, otherwise filterError will throw an error, because of
// the invalid CLG interface. All elements of values, except the error, if any,
// will not be included in the returned argument list.
func filterError(values []reflect.Value) ([]reflect.Value, error) {
	if len(values) == 0 {
		return nil, nil
	}

	var outputs []reflect.Value

	for i, v := range values {
		if err, ok := v.Interface().(error); ok {
			if i != len(values)-1 {
				// In golang we expect the error to be the last element of the output.
				// If this is not the case, we throw an error.
				return nil, maskAny(invalidInterfaceError)
			}
			if err != nil {
				// There was an error in the CLG output.
				return nil, maskAny(err)
			}
		}

		outputs = append(outputs, v)
	}

	return outputs, nil
}
