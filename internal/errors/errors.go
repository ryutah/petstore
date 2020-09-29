package errors

import (
	"errors"

	"github.com/hashicorp/go-multierror"
	"github.com/rotisserie/eris"
)

var (
	ErrServerError  = errors.New("server_error")
	ErrInvalidInput = errors.New("invalid_input")
	ErrNoSuchEntity = errors.New("no_such_entity")
)

var (
	Is     = eris.Is
	Wrap   = eris.Wrap
	Append = multierror.Append
)
