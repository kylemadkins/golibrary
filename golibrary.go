//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	id       int
	title    string
	chout    time.Time
	returned time.Time
}

type Library struct {
	books []Book
}

func displayCatalog(lib *Library) {
	for i := 0; i < len(lib.books); i++ {
		book := lib.books[i]
		fmt.Println("ID:", book.id)
		fmt.Println("Title:", book.title)
		fmt.Println("Checked Out:", book.chout)
		fmt.Println("Returned:", book.returned)
		fmt.Println("----------")
	}
}

func checkout(lib *Library, book *Book) {
	if time.Time.IsZero(book.chout) {
		book.chout = time.Now()
		book.returned = time.Time{}
		fmt.Println(book.title, "was checked out at", book.chout)
	} else {
		fmt.Println(book.title, "has not been returned")
	}
}

func ret(lib *Library, book *Book) {
	if !time.Time.IsZero(book.chout) {
		book.returned = time.Now()
		book.chout = time.Time{}
		fmt.Println(book.title, "was returned at", book.chout)
	} else {
		fmt.Println(book.title, "has not been checked out")
	}
}

func main() {
	tcitr := Book{id: 1, title: "The Catcher in the Rye"}
	tgg := Book{id: 2, title: "The Great Gatsby"}
	md := Book{id: 3, title: "Moby Dick"}
	lib := Library{[]Book{tcitr, tgg, md}}
	displayCatalog(&lib)
	checkout(&lib, &lib.books[0])
	checkout(&lib, &lib.books[0])
	ret(&lib, &lib.books[1])
	ret(&lib, &lib.books[0])
}
