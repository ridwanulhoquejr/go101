package main

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/go101/interfaces/email"
	"example.com/go101/interfaces/notification"
	"example.com/go101/interfaces/user"
)

func main() {

	// 1. Set up the real database
	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/myapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Build concrete implementations
	userRepo := user.NewRepository("trustmethisismydburl")
	mailClient := email.NewSMTPClient(
		"smtp.gmail.com",
		"noreply@myapp.com",
		"587",
	)

	// 3. Wire them into the notification sender
	//    Here's the magic: userRepo and mailClient
	//    SATISFY notification's interfaces without knowing it
	sender := notification.NewEmailSender(userRepo, mailClient)

	// 4. Use it
	if err := sender.SendWelcome("42"); err != nil {
		log.Fatalf("welcome failed: %v", err)
	}
	log.Println("welcome email sent!")

	//
	er := getError()
	if er != nil {
		fmt.Println("Boom!")

	}
	fmt.Println("not nill")

}

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func getError() error {
	// var err *MyError = nil
	// if err == nil {
	// 	return nil
	// }
	return nil
}
