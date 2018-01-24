package globber_test

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gentlemanautomaton/globber"
)

type Matcher interface {
	Match(string) bool
}

type Book struct {
	Title   string
	Authors []string
}

type BookSet []Book

func (books BookSet) ByTitle(matcher Matcher) (matched BookSet) {
	for _, book := range books {
		if matcher.Match(book.Title) {
			matched = append(matched, book)
		}
	}
	return
}

func (books BookSet) ByAuthor(matcher Matcher) (matched BookSet) {
	for _, book := range books {
		for _, author := range book.Authors {
			if matcher.Match(author) {
				matched = append(matched, book)
				break
			}
		}
	}
	return
}

var books = BookSet{
	// Freshman year
	{"Marmots 100: Orientation to marmotology", []string{"John Bramble", "T. Jones"}},
	{"Marmots 101: Native habitats and habits", []string{"John Bramble", "Grandpa George"}},
	{"Marmots 102: Typical predators in lakes", []string{"John Bramble", "T. Jones"}},
	// Sophomore year
	{"Marmots 200: Basic principals + methods", []string{"Smith", "Wesson", "McCloud"}},
	{"Marmots 201: Five 21st century policies", []string{"Don Juan the Iconoclast", "The Queen"}},
	{"Marmots 205: Marmot mascots and beyond?", []string{"Franky Frank Frankness", "Koala McGoo"}},
	// Comics
	{"No one messes with Marmotron or friends", []string{"John Dooey"}},
	{"How to spar with marmots, and feel okay", []string{"John Dooey", "Franky Dooey"}},
}

func Example() {
	// Defaults
	var (
		title  = globber.New("Marmots*")
		author = globber.NewSet("John Bramble", "T. Jones", "*Queen*")
	)

	// Inputs
	args := []string{"-title", "Marmots 20*", "-author", "*Bramble,*Franky*"}

	// Flag processing
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Var(&title, "title", "title search")
	fs.Var(&author, "author", "author search (separated by commas)")
	fs.Parse(args)

	// Title match by glob
	fmt.Printf("Books matching \"%s\" in the title:\n", title)
	for _, book := range books.ByTitle(title) {
		fmt.Printf("%s: %s\n", book.Title, strings.Join(book.Authors, ", "))
	}

	// Author match by glob set
	fmt.Printf("Books written by %s:\n", author)
	for _, book := range books.ByAuthor(author) {
		fmt.Printf("%s: %s\n", book.Title, strings.Join(book.Authors, ", "))
	}

	// Output:
	// Books matching "Marmots 20*" in the title:
	// Marmots 200: Basic principals + methods: Smith, Wesson, McCloud
	// Marmots 201: Five 21st century policies: Don Juan the Iconoclast, The Queen
	// Marmots 205: Marmot mascots and beyond?: Franky Frank Frankness, Koala McGoo
	// Books written by *Bramble,*Franky*:
	// Marmots 100: Orientation to marmotology: John Bramble, T. Jones
	// Marmots 101: Native habitats and habits: John Bramble, Grandpa George
	// Marmots 102: Typical predators in lakes: John Bramble, T. Jones
	// Marmots 205: Marmot mascots and beyond?: Franky Frank Frankness, Koala McGoo
	// How to spar with marmots, and feel okay: John Dooey, Franky Dooey
}
