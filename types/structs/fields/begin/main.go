// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
//

type author struct {
	name, surname string
}

func main() {
	// intialize author
	//
	muka := author{
		name:    "Muka",
		surname: "Lamejor de todos",
	}
	// print the author
	fmt.Printf("%#v\n", muka)
}
