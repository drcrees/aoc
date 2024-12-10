package p1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type File struct {
	id     int
	length int
	free   int
}

func Solve() {
	fmt.Println("--- 9-1 ---")

	file, err := os.Open("./2024/09/p1/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	files := parse(scanner.Text())
	strs := filesAsStrings(files)
	compact(strs)
	result := computeChecksum(strs)

	fmt.Printf("Result: %d\n", result)
}

func computeChecksum(strs []string) int {
	sum := 0
	for i := 0; i < len(strs); i++ {
		id, _ := strconv.Atoi(strs[i])
		sum += (i * id)
	}

	return sum
}

func compact(strs []string) {
	for i := len(strs) - 1; i >= 0; i-- {
		for j := 0; j < len(strs); j++ {
			if j == i {
				break
			}
			if strs[j] == "." {
				strs[j] = strs[i]
				strs[i] = "."
			}
		}
	}
}

func filesAsStrings(files []*File) []string {
	var strs []string
	for _, f := range files {
		for i := 0; i < f.length; i++ {
			strs = append(strs, strconv.Itoa(f.id))
		}
		for i := 0; i < f.free; i++ {
			strs = append(strs, ".")
		}
	}

	return strs
}

func printFiles(files []*File) {
	for _, f := range files {
		for i := 0; i < f.length; i++ {
			fmt.Printf("%d ", f.id)
		}
		for i := 0; i < f.free; i++ {
			fmt.Print(". ")
		}
	}
	fmt.Println()
}

func parse(diskMap string) (files []*File) {
	id := 0
	length := 0
	free := 0
	for i, r := range diskMap {
		if i%2 == 1 {
			free, _ = strconv.Atoi(string(r))
			files = append(files, &File{
				id:     id,
				length: length,
				free:   free,
			})
			id += 1
			length = 0
			free = 0
			continue
		}
		length, _ = strconv.Atoi(string(r))
		if i == len(diskMap)-1 {
			files = append(files, &File{
				id:     id,
				length: length,
				free:   free,
			})
		}
	}

	return files
}
