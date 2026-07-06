package core_http_types

import (
	"encoding/json"
	"reflect"

	"github.com/vasya2314/penny-plan/internal/core/domain"
	core_http_request "github.com/vasya2314/penny-plan/internal/core/transport/http/request"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func init() {
	// Teach the validator to unwrap Nullable[T] into the plain value it
	// holds (or nil, if the field wasn't present in the request), so that
	// regular tags like `min`, `max`, `startswith`, etc. can be reused as-is.
	//
	// Go generics turn Nullable[string], Nullable[int], etc. into distinct
	// types at runtime, so each concrete instantiation has to be listed
	// here - the same way the validator docs register one function for
	// sql.NullString, sql.NullInt64, sql.NullBool, etc. If a DTO ever needs
	// Nullable[T] for a T not listed below, add it to this call.
	core_http_request.RegisterCustomTypeFunc(
		unwrapNullable,
		Nullable[string]{},
		Nullable[int]{},
		Nullable[int64]{},
		Nullable[float64]{},
		Nullable[bool]{},
	)
}

// unwrapNullable extracts the underlying value out of a Nullable[T] field.
// It works for any T via reflection, since Go generics don't allow a single
// registration to cover every instantiation.
func unwrapNullable(field reflect.Value) any {
	if !field.FieldByName("Set").Bool() {
		return nil
	}

	return field.FieldByName("Value").Interface()
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true

	if string(b) == "null" {
		n.Value = nil

		return nil
	}

	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	n.Value = &value

	return nil
}

func (n *Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{
		Value: n.Value,
		Set:   n.Set,
	}
}
