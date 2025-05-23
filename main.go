package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	err := InitCommand()
	if err != nil {
		log.Fatal(err)
	}

	sampleMap := make(map[string]struct{})

	// Add sample to list
	for _, samplePath := range params.SamplePaths {
		err = AddFolderToLMap(samplePath, sampleMap)
		if err != nil {
			log.Fatalln("failed to add folder:", err)
		}
	}

	// check duplicated name
	dupList, err := CheckDuplicated(params.TargetPath, sampleMap)
	if err != nil {
		log.Fatalln("failed to check duplicate:", err)
	}

	log.Println("Duplicated: ", dupList)
	log.Println("Count: ", len(dupList))

	if !params.Confirm {
		fmt.Print("Delete duplicated？(y/N): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input != "y" && input != "yes" {
			log.Println("Canceled by user")
			return
		}
	}

	RemoveDuplicated(params.TargetPath, dupList)
}

func CheckDuplicated(targetPath string, sampleMap map[string]struct{}) ([]string, error) {
	var duplicated []string

	entries, err := os.ReadDir(targetPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if _, found := sampleMap[entry.Name()]; found {
			duplicated = append(duplicated, entry.Name())
		}
	}

	return duplicated, nil
}

func AddFolderToLMap(path string, sampleMap map[string]struct{}) error {
	// Read Folder
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// Add to list
	for _, entry := range entries {
		sampleMap[entry.Name()] = struct{}{}
	}

	return nil
}

func RemoveDuplicated(targetPath string, duplist []string) {

	for _, name := range duplist {
		currentPath := path.Join(targetPath, name)

		log.Println("Removing: ", currentPath)
		err := os.RemoveAll(currentPath)
		if err != nil {
			log.Println("Failed to remove: ", currentPath)
		}
	}
}
