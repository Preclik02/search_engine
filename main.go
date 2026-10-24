package main

import (

	//"strings"
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
	
		x := []string{entry.Name()}

		search_filename(x, path, clue)

		if entry.IsDir() {
			fmt.Println("Directory")  // -- DEBUG -- //
			fmt.Println(entry.Name()) // -- DEBUG -- // 
			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			directories = append(directories, pwd+"/"+entry.Name())
			//fmt.Println(directories)	// -- DEBUG -- //
		} else {				// -- DEBUG -- //
			fmt.Println("Not Directory")    // -- DEBUG -- // 
			fmt.Println(entry.Name()) 	// -- DEBUG -- //
		}
	
	}
	return directories
}

func search_filename(filename []string, path string, clue string) {

	j := 0
	for _, i := range filename {
		
		if i == filename[j] {
		
		}
		
		j ++
	}

}


func main() {

	var directories []string
	var path string
	var clue string
	var i int
	var size_of_directories int
	
	flag.StringVar(&clue, "clue", "", "string that the user remembers")
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path = pwd

	directories = check_and_add(directories, path, clue)
	for _, _ = range directories {
		size_of_directories ++
	}

	i = 0
	for k := 1; k <= size_of_directories; k ++{ 
		
		size_of_directories = 0
		for _, _ = range directories {
			size_of_directories ++ 
		}

		err := os.Chdir(directories[i])
		if err != nil {
			log.Fatal(err)
		}

		path = directories[i]

		directories = check_and_add(directories, path, clue)

		i++
	}

}
