package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TimRots/appmock/pkg/mock"
)

func main() {
	// get Mock struct from runtime name
	mock, err := mock.MockApp(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}

	// print stdout
	if mock.Stdout != "" {
		fmt.Println(mock.Stdout)
	}
	// print stderr
	if mock.Stderr != "" {
		fmt.Fprintf(os.Stderr, mock.Stderr)
	}
	// exit as desired
	os.Exit(mock.Exitstatus)
}
