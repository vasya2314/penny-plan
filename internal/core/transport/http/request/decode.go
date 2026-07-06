package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/vasya2314/penny-plan/internal/core/errors"
)

var requestValidator = validator.New()

// RegisterValidation adds a custom validation function under the given tag
// name, so it can be used in `validate:"..."` struct tags.
// See: https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Custom_Validation_Functions
func RegisterValidation(tag string, fn validator.Func) error {
	return requestValidator.RegisterValidation(tag, fn)
}

// RegisterCustomTypeFunc teaches the validator how to unwrap a custom type
// (e.g. a generic wrapper struct) into the plain value it holds, so that all
// the regular tags (required, min, max, startswith, ...) can be applied to it
// directly.
// See: https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Custom_Validation_Functions
func RegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...any) {
	requestValidator.RegisterCustomTypeFunc(fn, types...)
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf(
			"decode json %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	if err := requestValidator.Struct(dest); err != nil {
		return fmt.Errorf(
			"request validation %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
