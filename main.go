package main

import (

	"fmt"
	"os"
	"log"

)

func main() {

	var directories []string
	
	// -- read pwd -- //
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	
	// -- read entry by entry -- //
	for _, entry := range entries {

		// -- check && append if it is a directory -- //
		if entry.IsDir() {
			// -- debug -- //
			fmt.Println("Directory")
			fmt.Println(entry.Name())
			
			// -- append directory to slice -- //
			directories = append(directories, entry.Name())
		} else {
			fmt.Println("Not Directory")
			fmt.Println(entry.Name())
		}

	}
	// -- debug -- //
	fmt.Println(directories)

}
