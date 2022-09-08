package main

import "fmt"

// RUN - respoonsible for
// the initialitation and startup of the app
func Run() error {
	fmt.Println("starting up out application")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
