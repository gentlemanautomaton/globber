package globber

import (
	"strings"
	"unicode"
)

// Separator is the delimiter used when converting sets to and from strings.
const Separator = ','

// Set is a set of globs.
type Set []Glob

// NewSet returns a set of globs for the given pattern. Commas and
// whitespace will divide individual globs.
func NewSet(values ...string) (set Set) {
	for _, value := range values {
		if value != "" {
			set = append(set, New(value))
		}
	}
	return
}

// Split returns a set of globs for the given string. The string is split on
// commas and whitespace.
func Split(values string) (set Set) {
	sep := func(c rune) bool {
		return c == Separator || unicode.IsSpace(c)
	}
	for _, value := range strings.FieldsFunc(values, sep) {
		set = append(set, New(value))
	}
	return
}

// Match returns true if any member of the set matches the given value.
func (s Set) Match(value string) bool {
	for i := range s {
		if s[i].Match(value) {
			return true
		}
	}
	return false
}

// Set applies the given value or pattern to s. It assumes that values are
// delimited by whitespace or commas.
func (s *Set) Set(values string) error {
	*s = Split(values)
	return nil
}

// String returns a string representation of the set. Individual values are
// joined together by commas.
func (s Set) String() string {
	literals := make([]string, len(s))
	for i := range s {
		literals[i] = s[i].String()
	}
	return strings.Join(literals, string(Separator))
}
