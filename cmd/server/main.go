package main

import (
	"fmt"

	"github.com/sarapuertas/go-project/internal/comment"
	transportHttp "github.com/sarapuertas/go-project/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	} 
	
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
