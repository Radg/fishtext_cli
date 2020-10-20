package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(messageDefault)
		os.Exit(2)
	}

	command, count, err := ParseArguments(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	result, err := getContent(command, count)
	if err != nil {
		log.Fatal("Error occurred: ", err)
	}
	fmt.Println(result)
}
