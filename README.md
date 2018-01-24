# globber

[![GoDoc](https://godoc.org/github.com/gentlemanautomaton/globber?status.svg)](https://godoc.org/github.com/gentlemanautomaton/globber)

Package `globber` facilitates glob-style pattern matching for configuration.
It provides concrete types that are serializable as JSON and usable as flags
via `flag.Var()`

This is a small package that provides some convenient wrappers. The
hard work is performed by Sergey Kamardin's `glob` package, available here:
https://github.com/gobwas/glob

Please see documentation of the `glob` package for examples of patterns
and performance details.

Both `globber` and `glob` are MIT licensed.
