package p2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type File struct {
	id       int
	length   int
	free     int
	hasMoved bool
}

func Solve() {
	fmt.Println("--- 9-2 ---")

	file, err := os.Open("./2024/09/p2/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	files := parse(scanner.Text())
	compact(files)
	strs := filesAsStrings(files)
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

func compact(files []*File) {
	for i := len(files) - 1; i >= 0; i-- {
		for j := 0; j < len(files); j++ {
			if (files[j].free >= files[i].length) && !(j >= i) && !files[i].hasMoved {
				files[i].hasMoved = true
				files[i-1].free += (files[i].length + files[i].free)
				files[i].free = files[j].free - files[i].length
				files[j].free = 0
				file := files[i]
				newSlice := append(files[:i], files[i+1:]...)
				files = append(newSlice[:j+1], append([]*File{file}, newSlice[j+1:]...)...)
				i++
				break
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
