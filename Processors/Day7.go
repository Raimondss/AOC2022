package Processors

//Fuck this
import (
	"AOC_2022/FileReader"
	"fmt"
	"strconv"
	"strings"
)

type Day7 struct {
	Processor
	InputFilePath string
}

type File struct {
	name string
	size int
}

type Dir struct {
	name        string
	files       []File
	directories []Dir
	parentDir   *Dir
}

func (Dir Dir) getSize() int {
	totalSize := 0

	for _, file := range Dir.files {
		totalSize += file.size
	}

	for _, directory := range Dir.directories {
		totalSize += directory.getSize()
	}

	return totalSize
}

func filesByDirectory(inputArray []string) map[string][]File {
	fileSystem := make(map[string][]File)

	currentPath := ""
	totalFileSizes := 0
	for _, inputLine := range inputArray {
		lineElements := strings.Split(inputLine, " ")

		//Is number -> then file
		if fileSize, err := strconv.Atoi(lineElements[0]); err == nil {

			for _, dirFile := range fileSystem[currentPath] {
				if dirFile.name == lineElements[1] {
					continue
				}
			}

			totalFileSizes += fileSize
			file := File{name: lineElements[1], size: fileSize}
			fileSystem[currentPath] = append(fileSystem[currentPath], file)

			continue
		}

		if strings.Contains(inputLine, "$ cd") {
			if lineElements[2] == ".." {
				currentPathSplit := strings.Split(currentPath, "?")
				currentDirWithoutLast := currentPathSplit[:len(currentPathSplit)-1]

				currentPath = strings.Join(currentDirWithoutLast, "?")
				continue
			}

			if currentPath == "" {
				currentPath = lineElements[2]
				continue
			}

			currentPath = currentPath + "?" + lineElements[2]
			continue
		}

		continue

	}

	return fileSystem
}

func (Day7 Day7) FirstTaskResult() {
	fmt.Println("First task:")
	sizeByDirectory := getDirectoriesBySize(Day7)

	sizeSumOfDirectoriesOver100000 := 0
	for _, size := range sizeByDirectory {

		if size <= 100000 {
			sizeSumOfDirectoriesOver100000 += size
		}
	}

	fmt.Println(sizeSumOfDirectoriesOver100000)
}

func getDirectoriesBySize(Day7 Day7) map[string]int {
	inputArray := FileReader.GetStringArray(Day7.InputFilePath)

	sizeByDirectory := make(map[string]int)
	for directory, files := range filesByDirectory(inputArray) {
		subDirectories := strings.Split(directory, "?")

		fileSizes := 0
		for _, file := range files {
			fileSizes += file.size
		}

		for i, _ := range subDirectories {
			subDirectoryPath := strings.Join(subDirectories[:i+1], "?")
			sizeByDirectory[subDirectoryPath] += fileSizes
		}
	}

	return sizeByDirectory
}

func (Day7 Day7) SecondTaskResult() {
	fmt.Println("Second task:")

	directoriesBySize := getDirectoriesBySize(Day7)

	totalDiskSize := 70000000
	spaceNeeded := 30000000

	totalUsed := directoriesBySize["/"]
	currentSize := totalDiskSize - totalUsed

	spaceToFreeUp := spaceNeeded - currentSize

	var closest int
	for _, directorySize := range directoriesBySize {
		//Not valid for deletion
		if directorySize < spaceToFreeUp {
			continue
		}

		//first value
		if closest == 0 {
			closest = directorySize
			continue
		}

		if closest > directorySize {
			closest = directorySize
		}

	}

	println(closest)
}
