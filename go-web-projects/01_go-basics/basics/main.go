package main

import "log"

func main() {
	var whatToSSay string
	whatToSSay = saySomething("Life is good")

	var i int

	log.Println(whatToSSay)
	log.Println(i)

	var x string
	var y string

	x, y = demo("Hello")

	log.Println(x, y)

	x, _ = demo("Hello")

	log.Println(x)
}

func saySomething(s string) string {
	return s
}

func demo(s string) (string, string) {
	return s, "World"
}
