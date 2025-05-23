package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	targetPath := "D:\\#work\\BS\\CustomLevels"
	samplePath := "D:\\SteamLibrary\\steamapps\\common\\Beat Saber\\Levels\\【Fitness歌曲包】最新100首健身类型的高评分歌曲 更新至[2023-12-01]_WGzeyu"

	var targetNames []string
	var duplicatedPath []string
	targetMap := make(map[string]struct{})

	sampleEntries, err := os.ReadDir(samplePath)
	if err != nil {
		panic(err)
	}

	// 写入匹配样本
	for _, entry := range sampleEntries {
		targetNames = append(targetNames, entry.Name())
	}

	for _, name := range targetNames {
		targetMap[name] = struct{}{}
	}

	entries, err := os.ReadDir(targetPath)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if _, found := targetMap[entry.Name()]; found {
			fmt.Println("Found duplicated: ", entry.Name())
			duplicatedPath = append(duplicatedPath, entry.Name())
		}
	}

	fmt.Println("result: ", duplicatedPath)
	fmt.Println("count: ", len(duplicatedPath))

	var re string
	fmt.Scan(&re)

	if re != "" {
		for _, name := range duplicatedPath {
			//fmt.Println(path.Join(targetPath, name))
			os.RemoveAll(path.Join(targetPath, name))
		}
	}
}
