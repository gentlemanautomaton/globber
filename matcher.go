package globber

// Matcher is a generic matching interface satisfied by globs and sets of globs.
type Matcher interface {
	Match(string) bool
}
