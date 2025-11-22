package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	var parking *ParkingLot = nil

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input := strings.Split(line, " ")

		parking, err = Command(parking, input)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
