package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	targetPath := "D:\\#work\\BS\\CustomLevels"
	samplePath := "D:\\SteamLibrary\\steamapps\\common\\Beat Saber\\Levels\\【Fitness歌曲包】最新100首健身类型的高评分歌曲 更新至[2023-12-01]_WGzeyu"

	sampleMap := make(map[string]struct{})

	// Add sample to list
	err := AddFolderToLMap(samplePath, sampleMap)
	if err != nil {
		log.Fatalln("failed to add folder:", err)
	}

	// check duplicated name
	dupList, err := CheckDuplicated(targetPath, sampleMap)
	if err != nil {
		log.Fatalln("failed to check duplicate:", err)
	}

	log.Println("result: ", dupList)
	log.Println("count: ", len(dupList))

	var re string
	fmt.Scan(&re)

	if re != "" {
		RemoveDuplicated(targetPath, dupList)
	}
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
