package main

import (
	"fmt"
	"context"
    
	"github.com/sarapuertas/go-project/internal/db"
)

// RUN - respoonsible for
// the initialitation and startup of the app
func Run() error {
	fmt.Println("starting up out application")

    db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
    if err := db.Ping(context.Background()); err != nil {
		return err
	} 
    fmt.Println("successfully connected and pinged to the DB!")
	
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
