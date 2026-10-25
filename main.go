package main

import (

	"flag"
	"fmt"
	"os"
	"log"

)

func check_and_add(directories []string, path string, clue_slice []string) []string {
	
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, entry := range entries {
	
		x := []string{entry.Name()}
		var y []string

		for _, character := range x[0] {
			y = append(y, string(character))
		}

		search_filename(entry.Name(), y, path, clue_slice)

		if entry.IsDir() {
			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			directories = append(directories, pwd+"/"+entry.Name())
		}
	
	}
	return directories
}

func search_filename(filename_string string, filename_slice []string, path string, clue_slice []string) {

	size_of_filename := 0
	for _, _ = range filename_slice {
		size_of_filename ++
	}
	
	var characters_right int

	//fmt.Println(clue_slice)
	//fmt.Println(filename_slice)
	//fmt.Println(filename_string)

	for _, i := range filename_slice {
		for _, j := range clue_slice {	
			if i == j {
				characters_right += 1
			}
		}
		
	}

	if characters_right >= 1 {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pwd+"/"+filename_string)
	}
	
}


func main() {

	var directories []string
	var path string
	var clue string
	var i int
	var size_of_directories int
	var clue_slice []string

	flag.StringVar(&clue, "clue", "", "string that the user remembers")
	flag.Parse()

	for _, character := range clue {
		clue_slice = append(clue_slice, string(character))
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path = pwd

	directories = check_and_add(directories, path, clue_slice)
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

		directories = check_and_add(directories, path, clue_slice)

		i++
	}

}
