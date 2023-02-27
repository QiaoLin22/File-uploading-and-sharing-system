package main

import (
	"context"
	"fmt"

	"github.com/QiaoLin22/Golang-production-1/internal/comment"
	"github.com/QiaoLin22/Golang-production-1/internal/db"
)

// initialize and start up of the go application
func Run() error {
	fmt.Println("starting up application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	fmt.Println("Connected to and pinged database")

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"71c5d074-b6cf-11ec-b909-0242ac120002",
	))
	return nil
}

func main() {
	fmt.Println("Hello")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
