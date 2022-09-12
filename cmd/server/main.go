package main

import (
	"fmt"
    "context"
	"github.com/sarapuertas/go-project/internal/comment"
    
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
    if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate db")
		return err
	} 
	

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID: "50a174fa-3296-11ed-a261-0242ac120002",
			Slug: "manual-test",
			Author: "Sarita",
			Body: "Hello Woooooorld :p",
		},
	)

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"50a174fa-3296-11ed-a261-0242ac120002",
	))
	
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
