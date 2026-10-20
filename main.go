package main

import (

	"fmt"
	"os"
	"log"

)

func check_and_add(directories []string, path string) []string {
	
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, entry := range entries {
		
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

func main() {

	var directories []string
	var path string

	path = "./"

	directories = check_and_add(directories, path)

	fmt.Println(directories)

}
