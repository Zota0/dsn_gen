package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Define ANSI escape codes for colors
	const (
		Reset        = "\033[0m"
		Bold         = "\033[1m"
		Italic       = "\033[3m"
		Underline    = "\033[4m"
		ReverseVideo = "\033[7m"
		Red          = "\033[31m"
		Green        = "\033[32m"
		Yellow       = "\033[33m"
		Blue         = "\033[34m"
		Magenta      = "\033[35m"
		Cyan         = "\033[36m"
		White        = "\033[37m"
	)

	fmt.Printf("%s", Reset)

	// Helper function to read input and trim it
	readInput := func(prompt string, color string) string {
		fmt.Print(Bold + color + prompt + Reset + Yellow)
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}

	host := readInput("Host: \n >\t", Cyan)
	port := readInput("Port: \n >\t", Magenta)
	db := readInput("Database: \n >\t", Cyan)
	user := readInput("Username: \n >\t", Magenta)
	pass := readInput("Password: \n >\t", Cyan) // Optional: hide password input
	ssl := readInput("SSL Mode: \n >\t", Magenta)

	if ssl == "" {
		ssl = "disable"
	}

	// Printing the DSN
	fmt.Print("\n" + Bold + "DSN: \n" + Reset)
	fmt.Printf("%s%shost=%s;port=%s;user=%s;password=%s;dbname=%s;sslmode=%s", White, Italic, host, port, user, pass, db, ssl)
	fmt.Printf("\n\n%s", Reset)
}
