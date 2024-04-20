package main

import "os"

func writefile(tag string, notes []string) error {
	filename := tag + ".txt"
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, note := range notes {
		_, err := file.WriteString(note + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
