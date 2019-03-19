package globber

import (
	"encoding/json"
	"strings"

	"github.com/gobwas/glob"
)

// Glob matches string literals and glob patterns.
type Glob struct {
	literal string
	exact   glob.Glob
	lowered glob.Glob
}

// New returns a glob for the given string literal or pattern.
//
// If value cannot be compiled as a pattern, the result will only match
// literal values.
func New(value string) Glob {
	exact, _ := glob.Compile(value)
	lowered, _ := glob.Compile(strings.ToLower(value))
	return Glob{
		literal: value,
		exact:   exact,
		lowered: lowered,
	}
}

// Match returns true if the glob matches the given value exactly or as a
// pattern match.
func (g Glob) Match(value string) bool {
	if g.literal == value {
		return true
	}
	if g.exact != nil {
		return g.exact.Match(value)
	}
	return false
}

// MatchInsensitive returns true if the glob case-insensitively matches the
// given value exactly or as a pattern match.
func (g Glob) MatchInsensitive(value string) bool {
	if strings.EqualFold(g.literal, value) {
		return true
	}
	if g.lowered != nil {
		return g.lowered.Match(strings.ToLower(value))
	}
	return false
}

// Set applies the given value or pattern to g. It facilitates use in the flag
// package.
func (g *Glob) Set(value string) error {
	g.literal = value
	g.exact, _ = glob.Compile(value)
	g.lowered, _ = glob.Compile(strings.ToLower(value))
	return nil
}

// String returns a string representation of the value or pattern.
func (g Glob) String() string {
	return g.literal
}

// MarshalJSON marshals the glob as a JSON-encoded string.
func (g Glob) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.literal)
}

// UnmarshalJSON unmarshals the glob from a JSON-encoded string.
func (g *Glob) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &g.literal); err != nil {
		return err
	}
	g.exact, _ = glob.Compile(g.literal)
	g.lowered, _ = glob.Compile(strings.ToLower(g.literal))
	return nil
}
