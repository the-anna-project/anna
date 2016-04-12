package clg

import (
	"math"
)

func (i *clgIndex) DivideFloat64(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 / i2

	return []interface{}{s}, nil
}

func (i *clgIndex) GreaterThanFloat64(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	greater := i1 > i2

	return []interface{}{greater}, nil
}

func (i *clgIndex) LesserThanFloat64(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	lesser := i1 < i2

	return []interface{}{lesser}, nil
}

func (i *clgIndex) MultiplyFloat64(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 * i2

	return []interface{}{s}, nil
}

func (i *clgIndex) PowFloat64(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := math.Pow(i1, i2)

	return []interface{}{s}, nil
}

func (i *clgIndex) SqrtFloat64(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	s := math.Sqrt(i1)

	return []interface{}{s}, nil
}

func (i *clgIndex) SubtractFloat64(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 - i2

	return []interface{}{s}, nil
}

func (i *clgIndex) SumFloat64(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 + i2

	return []interface{}{s}, nil
}
