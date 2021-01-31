package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	mySlice := []string{"dog", "cat", "mouse"}

	for j, x := range mySlice {
		log.Println(j, x)
	}

	for _, y := range mySlice {
		log.Println(y)
	}

	myMap := make(map[string]string)
	myMap["hello"] = "hola"
	myMap["bye"] = "adios"

	for key, value := range myMap {
		log.Println(key, value)
	}
}
