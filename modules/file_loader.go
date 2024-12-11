package modules

import (
	"bufio"
	"os"
	"log"
	"fmt"
)

func isEmpty(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err // erro on acess file 
	}
	return info.Size() == 0, nil
}

func Loader(wordlistPath string) []string {
	file, err := os.Open(wordlistPath)
	if err != nil {
		log.Fatal("File does not exist or invalid path!")
	}
	defer file.Close()

	// reading wordlist from file
	var Wordlist []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Wordlist = append(Wordlist, scanner.Text()) 
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Scanning error!")
	}
	// checking if file is empty	
	empty, err := isEmpty(wordlistPath)
	if err != nil {
		fmt.Printf("Error checking file: %v\n", err)
	}
	if empty {
		fmt.Println("Wordlist is empty!")
	}

	return Wordlist
}
