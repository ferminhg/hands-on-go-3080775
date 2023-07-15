// challenges/types-composite/begin/main.go
package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author

type book struct {
	title  string
	author author
}

// define a library type to track a list of books
//
//	type library struct {
//		list
//	}
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(a author) []book {
	return l[a.name]
}

func main() {
	// create a new library
	l := library{}

	// add 2 books to the library by the same author
	l.addBook(book{
		title:  "wop",
		author: author{name: "Muka"},
	})

	l.addBook(book{
		title:  "wop wop",
		author: author{name: "Muka"},
	})

	// add 1 book to the library by a different author
	l.addBook(book{
		title:  "wiiii",
		author: author{name: "Beka"},
	})

	// dump the library with spew
	spew.Dump(l)

	// lookup books by known author in the library
	spew.Println(l.lookupByAuthor(author{name: "Muka"}))

	// print out the first book's title and its author's name }
	books := l.lookupByAuthor(author{name: "Muka"})
	if len(books) != 0 {
		fmt.Printf("%v by %s\n", books[0].author, books[0].title)
	}
	//spew.Dump(l.lookupByAuthor(author{name: "Muka"})[0])
}
