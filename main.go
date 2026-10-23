package main

import (

	"flag"
	"fmt"
	"os"
	"log"

)

func check_and_add(directories []string, path string, clue string) []string {
	
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, entry := range entries {
	
		search_filename(entry.Name(), path, clue)

		if entry.IsDir() {
			fmt.Println("Directory")
			fmt.Println(entry.Name())
			directories = append(directories, entry.Name())
			fmt.Println(directories)
		} else {
			fmt.Println("Not Directory")
			fmt.Println(entry.Name())
		}
	
	}
	return directories
}

func search_filename(filename string, path string, clue string) {

	// -- seaching logic -- //

}


func main() {

	var directories []string
	var path string
	var clue string
	var i int
	
	flag.StringVar(&clue, "clue", "", "string that the user remembers")

	flag.Parse()

	fmt.Println(clue)

	path = "./"

	i = 0
	for _, _ = range directories { 
		
		directories = check_and_add(directories, path, clue)

		fmt.Println(directories[i])
		

		i++
	}

}
