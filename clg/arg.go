package clg

func ArgToBool(args []interface{}, index int) (bool, error) {
	if len(args) < index+1 {
		return false, maskAnyf(notEnoughArgumentsError, "expected %d got %d", index+1, len(args))
	}

	if s, ok := args[index].(bool); ok {
		return s, nil
	}

	return false, maskAnyf(wrongArgumentTypeError, "expected %T got %T", "", args[index])
}

func ArgToString(args []interface{}, index int) (string, error) {
	if len(args) < index+1 {
		return "", maskAnyf(notEnoughArgumentsError, "expected %d got %d", index+1, len(args))
	}

	if s, ok := args[index].(string); ok {
		return s, nil
	}

	return "", maskAnyf(wrongArgumentTypeError, "expected %T got %T", "", args[index])
}

func ArgToStringSlice(args []interface{}, index int) ([]string, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d got %d", index+1, len(args))
	}

	if ss, ok := args[index].([]string); ok {
		return ss, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected %T got %T", "", args[index])
}
