package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("--- 11-2 ---")

	file, err := os.Open("./input2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result = 0

	for scanner.Scan() {
		parse(scanner.Text())
	}

	fmt.Printf("Result: %v", result)
}

func parse(str string) {
	pixelValues := strings.Split(str, "")

	var pixels []Pixel
	for _, p := range pixelValues {
		pixels = append(pixels, Pixel{value: p, isGalaxy: p == "#"})
	}
	image = append(image, pixels)
}
