package main

import (
	"log"

	"github.com/SPahooja/QoHash/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatal("Something went wrong\n", err)
	}
}
