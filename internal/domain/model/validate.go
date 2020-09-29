package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ryutah/petstore/internal/errors"
)

func validate(v interface{}) error {
	err := validator.New().Struct(v)
	if err == nil {
		return nil
	}

	var errs error
	ve := err.(validator.ValidationErrors)
	for _, e := range ve {
		errs = errors.Append(errs, errors.Wrap(errors.ErrInvalidInput, fmt.Sprint(e)))
	}
	return errs
}
