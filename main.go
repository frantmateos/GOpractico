package main 

import "fmt"

func main (){ 
	fmt.Println("Ingresa una URL: ")

	var input string 

	fmt.Scanln(&input)

	result := shortURL(input) //: declara variable y = asigna vsalor 

	fmt.Println("shortURL: ",result)
}

func shortURL(input string) string {
	return fmt.Sprintf("hhtp://shorturl.com/abc123")
}