package main

import (
	"os"
	"fmt"
	"flag"
	"log"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No commands provided")
		os.Exit(2)
	}
	fmt.Println(os.Args )
	cmd := os.Args[1]

	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "CLI BASICS - REMINDER CLI", "The message to greet command")
		err := greetCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Hello and welcome %s\n", *msgFlag)
	case "help":
		fmt.Println("Some help message")
	default: 
		fmt.Printf("Unknown command: %s\n", cmd)
	}
}