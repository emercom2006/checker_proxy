package main

import (
	"bufio"
	"os"
)

func ReadFile(nameFileText string) map[int]string {
	mapProxy := make(map[int]string)
	file, _ := os.Open(nameFileText)
	scanner := bufio.NewScanner(file)
	i := 0
	defer file.Close()
	for scanner.Scan() {
		mapProxy[i] = scanner.Text()
		i++
	}
	return mapProxy
}
