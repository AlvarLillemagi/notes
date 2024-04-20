package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 || (len(os.Args) == 2 && os.Args[1] == "help") {
		displayHelpMessage()
		return
	}

	tag := os.Args[1]
	fmt.Println("\nWelcome to the notes tool for", tag)

	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			err := displayNotes(tag)
			if err != nil {
				fmt.Println("\nError displaying notes:", err)
			}
		case 2:
			err := addNote(tag)
			if err != nil {
				fmt.Println("\nError adding note:", err)
			}
		case 3:
			err := deleteNote(tag)
			if err != nil {
				fmt.Println("\nError deleting note:", err)
			}
		case 4:
			fmt.Println("\nExiting...")
			return
		default:
			fmt.Println("\nInvalid choice. Please choose again.")
		}
	}
}

func displayHelpMessage() {
	fmt.Println("Tool functionality:")
	fmt.Println("1. Display notes from the collection")
	fmt.Println("2. Add a note to the collection")
	fmt.Println("3. Remove a note from the collection")
	fmt.Println("4. Exit the program")
}

func displayNotes(tag string) error {
	notes, err := readfile(tag)
	if err != nil {
		return err
	}
	fmt.Println("\nNotes:")
	for i, note := range notes {
		fmt.Printf("%03d - %s\n", i+1, note)
	}
	fmt.Println()
	return nil
}

func addNote(tag string) error {
	fmt.Println("\nEnter the note text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	note := scanner.Text()

	notes, err := readfile(tag)
	if err != nil {
		return err
	}
	notes = append(notes, note)

	err = writefile(tag, notes)
	if err != nil {
		return err
	}

	fmt.Println("\nNote added successfully.")
	return nil
}

func deleteNote(tag string) error {
	fmt.Println("\nEnter the number of the note to remove or 0 to cancel:")
	var input string
	fmt.Scanln(&input)
	noteNum, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	if noteNum == 0 {
		fmt.Println("\nOperation canceled.")
		return nil
	}

	notes, err := readfile(tag)
	if err != nil {
		return err
	}

	if noteNum < 1 || noteNum > len(notes) {
		fmt.Println("\nInvalid note number.")
		return nil
	}

	notes = append(notes[:noteNum-1], notes[noteNum:]...)

	err = writefile(tag, notes)
	if err != nil {
		return err
	}

	fmt.Println("\nNote removed successfully.")
	return nil
}
