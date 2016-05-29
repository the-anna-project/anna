package text

import (
	"fmt"

	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func maskAnyf(err error, f string, v ...interface{}) error {
	if err == nil {
		return nil
	}

	f = fmt.Sprintf("%s: %s", err.Error(), f)
	newErr := errgo.WithCausef(nil, errgo.Cause(err), f, v...)
	newErr.(*errgo.Err).SetLocation(1)

	return newErr
}

var invalidConfigError = errgo.New("invalid config")

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return errgo.Cause(err) == invalidConfigError
}

var gatewayClosedError = errgo.New("gateway closed")

// IsGatewayClosed asserts gatewayClosedError.
func IsGatewayClosed(err error) bool {
	return errgo.Cause(err) == gatewayClosedError
}

var invalidRequestError = errgo.New("invalid request")

// IsInvalidRequest asserts invalidRequestError.
func IsInvalidRequest(err error) bool {
	return errgo.Cause(err) == invalidRequestError
}
