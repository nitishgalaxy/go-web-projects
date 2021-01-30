package main

import (
	"log"

	"github.com/nitishgalaxy/myniceprogram/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)
}
