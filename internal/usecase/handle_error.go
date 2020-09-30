package usecase

import (
	"context"

	"github.com/ryutah/petstore/internal/errors"
)

type errorOutputPort interface {
	InvalidInput(context.Context)
	NotFound(context.Context)
	ServerError(context.Context)
}

type errorHandlerFunc func(context.Context, error, errorOutputPort)

type errorHandlers struct {
	serverError  errorHandlerFunc
	invalidInput errorHandlerFunc
	noSuchEntity errorHandlerFunc
}

func newErrorHandlers(opts ...errrorHandlerOption) *errorHandlers {
	handler := errorHandlers{
		serverError: func(ctx context.Context, err error, out errorOutputPort) {
			out.ServerError(ctx)
		},
		invalidInput: func(ctx context.Context, err error, out errorOutputPort) {
			out.InvalidInput(ctx)
		},
		noSuchEntity: func(ctx context.Context, err error, out errorOutputPort) {
			out.NotFound(ctx)
		},
	}
	for _, opt := range opts {
		opt(&handler)
	}
	return &handler
}

type errrorHandlerOption func(*errorHandlers)

func withNoSuchEntityFunc(f errorHandlerFunc) errrorHandlerOption {
	return func(e *errorHandlers) {
		e.noSuchEntity = f
	}
}

// always return false
func handleError(ctx context.Context, err error, out errorOutputPort, opts ...errrorHandlerOption) bool {
	handlers := newErrorHandlers(opts...)

	switch {
	case errors.Is(err, errors.ErrServerError):
		handlers.serverError(ctx, err, out)
	case errors.Is(err, errors.ErrInvalidInput):
		handlers.invalidInput(ctx, err, out)
	case errors.Is(err, errors.ErrNoSuchEntity):
		handlers.noSuchEntity(ctx, err, out)
	default:
		handlers.serverError(ctx, err, out)
	}

	return false
}
