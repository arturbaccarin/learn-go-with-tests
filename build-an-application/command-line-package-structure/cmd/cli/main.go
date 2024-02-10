package main

import "fmt"

func main() {
	/*
		os.Stdin is what we'll use in main to capture the user's input.
		It is a *File under the hood which means it implements io.Reader
		which as we know by now is a handy way of capturing text
	*/
	fmt.Println("Let's play poker")
}
