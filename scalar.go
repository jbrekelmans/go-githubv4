package githubv4

import (
	"encoding"
	"fmt"
	"math/big"
	"time"

	"github.com/jbrekelmans/go-graphql"
)

// This file defines GitHub GraphQL scalars.
// See https://docs.github.com/en/graphql/reference/scalars#about-scalars
//
// github.com/jbrekelmans/go-graphql uses reflect to declare variables in GraphQL.
// The Go struct name is used as the GraphQL type name, so we need to define
// scalars as structs.
// GraphQL scalars like Boolean, Float and Int are mapped to bool, float64 and int
// types, and need not be defined as a struct.

// Base64String represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as Base64String in GraphQL as expected by the GitHub API.
//
// Base64String is a (potentially binary) string encoded using base64.
type Base64String struct {
	S string
}

var _ encoding.TextMarshaler = (*Base64String)(nil)
var _ encoding.TextUnmarshaler = (*Base64String)(nil)

func (base64String Base64String) MarshalText() ([]byte, error) {
	return []byte(base64String.S), nil
}

func (base64String *Base64String) UnmarshalText(b []byte) error {
	base64String.S = string(b)
	return nil
}

// BigInt represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as BigInt in GraphQL as expected by the GitHub API.
//
// BigInt represents non-fractional signed whole numeric values. Since the value may exceed the
// size of a 32-bit integer, it's encoded as a string.
type BigInt struct {
	N big.Int
}

var _ encoding.TextMarshaler = (*BigInt)(nil)
var _ encoding.TextUnmarshaler = (*BigInt)(nil)

func (bigInt BigInt) MarshalText() ([]byte, error) {
	s := bigInt.N.Text(10)
	return []byte(s), nil
}

func (bigInt *BigInt) UnmarshalText(b []byte) error {
	s := string(b)
	if _, ok := bigInt.N.SetString(s, 10); !ok {
		return fmt.Errorf(`error in (*BigInt).UnmarshalText: could not parse string as base10 integer: %#v`, s)
	}
	return nil
}

// Date represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as Date in GraphQL as expected by the GitHub API.
//
// Date is an ISO-8601 encoded date string.
type Date struct {
	S string
}

var _ encoding.TextMarshaler = (*Date)(nil)
var _ encoding.TextUnmarshaler = (*Date)(nil)

func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.S), nil
}

func (d *Date) UnmarshalText(b []byte) error {
	d.S = string(b)
	return nil
}

// DateTime represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as DateTime in GraphQL as expected by the GitHub API.
//
// DateTime is an ISO-8601 encoded UTC date string.
type DateTime struct {
	time.Time

	// DateTime inherits the method set from time.Time. In particular, MarshalText is inherited.
	// (time.Time).MarshalText sends subsecond precision to GitHub.
	// If this causes errors on the GitHub side, we may want to define our own
	// MarshalText that errors early.
}

// GitObjectID represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as GitObjectID in GraphQL as expected by the GitHub API.
//
// GitObjectID is a Git object ID.
type GitObjectID struct {
	S string
}

var _ encoding.TextMarshaler = (*Date)(nil)
var _ encoding.TextUnmarshaler = (*Date)(nil)

func (g GitObjectID) MarshalText() ([]byte, error) {
	return []byte(g.S), nil
}

func (g *GitObjectID) UnmarshalText(b []byte) error {
	g.S = string(b)
	return nil
}

// GitTimestamp represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as GitTimestamp in GraphQL as expected by the GitHub API.
//
// GitTimestamp is an ISO-8601 encoded date string. Unlike the DateTime type, GitTimestamp is not converted in UTC.
type GitTimestamp struct {
	S string
}

var _ encoding.TextMarshaler = (*Date)(nil)
var _ encoding.TextUnmarshaler = (*Date)(nil)

func (g GitTimestamp) MarshalText() ([]byte, error) {
	return []byte(g.S), nil
}

func (g *GitTimestamp) UnmarshalText(b []byte) error {
	g.S = string(b)
	return nil
}

// TODO add GitRefname
// GitRefname is currently in preview:
// https://docs.github.com/en/graphql/reference/scalars#gitrefname

// GitSSHRemote represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as GitSSHRemote in GraphQL as expected by the GitHub API.
//
// GitSSHRemote is a git SSH string.
type GitSSHRemote struct {
	S string
}

var _ encoding.TextMarshaler = (*Date)(nil)
var _ encoding.TextUnmarshaler = (*Date)(nil)

func (g GitSSHRemote) MarshalText() ([]byte, error) {
	return []byte(g.S), nil
}

func (g *GitSSHRemote) UnmarshalText(b []byte) error {
	g.S = string(b)
	return nil
}

// HTML represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as HTML in GraphQL as expected by the GitHub API.
//
// HTML is a string containing HTML code.
type HTML struct {
	S string
}

var _ encoding.TextMarshaler = (*Date)(nil)
var _ encoding.TextUnmarshaler = (*Date)(nil)

func (h HTML) MarshalText() ([]byte, error) {
	return []byte(h.S), nil
}

func (h *HTML) UnmarshalText(b []byte) error {
	h.S = string(b)
	return nil
}

// ID represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as ID in GraphQL as expected by the GitHub API.
//
// ID represents a unique identifier that is Base64 obfuscated. It
// is often used to refetch an object or as key for a cache. The ID
// type appears in a JSON response as a String; however, it is not
// intended to be human-readable. When expected as an input type,
// any string (such as "VXNlci0xMA==") or integer (such as 4) input
// value will be accepted as an ID.
type ID = graphql.ID

// PreciseDateTime represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as PreciseDateTime in GraphQL as expected by the GitHub API.
//
// An ISO-8601 encoded UTC date string with millisecond precision.
type PreciseDateTime struct {
	time.Time

	// PreciseDateTime inherits the method set from time.Time. In particular, MarshalText is inherited.
	// (time.Time).MarshalText sends nanosecond precision to GitHub.
	// If this causes errors on the GitHub side, we may want to define our own
	// MarshalText that errors early.
}

// URI represents a GraphQL scalar. Use as the type for variables values in queries/mutations
// to ensure the variable is declared as URI in GraphQL as expected by the GitHub API.
//
// URI is an RFC 3986, RFC 3987, and RFC 6570 (level 4) compliant URI string.
type URI struct {
	S string
}

var _ encoding.TextMarshaler = (*Date)(nil)
var _ encoding.TextUnmarshaler = (*Date)(nil)

func (u URI) MarshalText() ([]byte, error) {
	return []byte(u.S), nil
}

func (u *URI) UnmarshalText(b []byte) error {
	u.S = string(b)
	return nil
}
