package githubv4

import (
	"encoding/json"

	"github.com/jbrekelmans/go-graphql"
)

// enhanceError parses GitHub-specific entries of response errors, that are not
// specified in the GraphQL specification, for convenient access.
// enhanceError replaces a *graphql.Error with a *Error.
// If err is not a *graphql.Error then returns err unmodified.
func enhanceError(err error) error {
	base, ok := err.(*graphql.Error)
	if !ok {
		return err
	}
	enhanced := &Error{
		Err:     base.Err,
		Message: base.Message,
	}
	if base.Errors != nil {
		enhanced.Errors = make([]ErrorItem, 0, len(base.Errors))
		for _, baseItem := range base.Errors {
			enhanced.Errors = append(enhanced.Errors, enhanceErrorItem(baseItem))
		}
	}
	return enhanced
}

func enhanceErrorItem(base graphql.ErrorItem) ErrorItem {
	return ErrorItem{
		Extensions: tryGet[map[string]any](base.Raw, "extensions"),
		Message:    base.Message,
		Type:       tryGet[string](base.Raw, "type"),
		Raw:        base.Raw,
	}
}

// Error is an error type used by *Client to feed back GraphQL-level errors.
type Error struct {
	// Err is the wrapped error.
	Err error

	// Errors are the response errors.
	// Errors reflects the "errors" property of JSON objects in HTTP response bodies.
	// See https://spec.graphql.org/.
	Errors []ErrorItem

	Message string
}

var _ error = (*Error)(nil)

// Error implements the error interface.
func (e *Error) Error() string {
	return e.Message
}

// Unwrap supports Golang 1.13+ error wrapping. See https://go.dev/blog/go1.13-errors
func (e *Error) Unwrap() error {
	return e.Err
}

// ErrorItem is a response error. See https://spec.graphql.org/.
type ErrorItem struct {
	// Error message.
	Message string

	// Value of "extensions" entry.
	// See https://spec.graphql.org/.
	Extensions map[string]any

	// Raw entries as per the JSON value returned by the server.
	// Does not include entries that were successfully parsed into
	// their corresponding fields.
	Raw map[string]json.RawMessage

	// Value of "type" entry.
	// Although this entry is unspecified in the GraphQL specification https://spec.graphql.org/,
	// GitHub sets this for some errors.
	// Known values (as seen in the wild): RATE_LIMITED
	Type string
}

func tryGet[T any](raw map[string]json.RawMessage, key string) T {
	var valueZero T
	valueJSON, ok := raw[key]
	if !ok {
		return valueZero
	}
	var valueParsed T
	if err := json.Unmarshal(valueJSON, &valueParsed); err != nil {
		return valueZero
	}
	delete(raw, key)
	return valueParsed
}
