package globber

import (
	"encoding/json"

	"github.com/gobwas/glob"
)

// Glob matches string literals and glob patterns.
type Glob struct {
	literal string
	glob    glob.Glob
}

// New returns a glob for the given string literal or pattern.
//
// If value cannot be compiled as a pattern, the result will only match
// literal values.
func New(value string) Glob {
	g, _ := glob.Compile(value)
	return Glob{
		literal: value,
		glob:    g,
	}
}

// Match returns true if the glob matches the given value exactly or as a
// pattern match.
func (g Glob) Match(value string) bool {
	if g.literal == value {
		return true
	}
	if g.glob != nil {
		return g.glob.Match(value)
	}
	return false
}

// Set applies the given value or pattern to g. It facilities use in the flag
// package.
func (g *Glob) Set(value string) error {
	g.literal = value
	g.glob, _ = glob.Compile(value)
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
	g.glob, _ = glob.Compile(g.literal)
	return nil
}
