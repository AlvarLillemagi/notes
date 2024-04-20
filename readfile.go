package main

import (
	"bufio"
	"os"
)

func readfile(tag string) ([]string, error) {
	filename := tag + ".txt"
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
