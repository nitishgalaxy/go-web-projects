package main

import (
	"github.com/tsawler/myniceprogram/helpers"
	"log"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)
}
