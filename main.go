package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Define color functions
	cyan := color.New(color.FgCyan).Add(color.Bold)
	magenta := color.New(color.FgMagenta).Add(color.Bold)
	yellow := color.New(color.FgYellow).Add(color.Bold)
	green := color.New(color.FgGreen).Add(color.Bold)
	white := color.New(color.FgWhite).Add(color.Bold)
	blue := color.New(color.FgBlue).Add(color.Bold)

	// Helper function to read input and trim it
	readInput := func(prompt string, promptColor *color.Color, inputColor *color.Color) string {
		promptColor.Print(prompt)
		inputColor.Print("> ")
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}

	// Function to display a selection menu
	displayMenu := func(prompt string, options []string) string {
		cyan.Println(prompt)
		colors := []*color.Color{magenta, blue, green, white}
		for i, option := range options {
			colors[i].Printf("%d. %s\n", i+1, option)
		}
		yellow.Print("> ")
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}

	// Display the selection menu for database type
	dbType := displayMenu("Choose database type:", []string{"Postgres", "MySQL", "SQLite", "Other"})

	// Get common DSN parameters
	host := readInput("Host: \n", cyan, yellow)
	port := readInput("Port: \n", magenta, green)
	db := readInput("Database: \n", cyan, yellow)
	user := readInput("Username: \n", magenta, green)
	pass := readInput("Password: \n", cyan, yellow)
	ssl := readInput("SSL Mode: \n", magenta, green)
	if ssl == "" {
		ssl = "disable"
	}

	// Construct the DSN based on the selected database type
	var dsn string
	switch dbType {
	case "1": // Postgres
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, pass, host, port, db, ssl)
	case "2": // MySQL
		dsn = fmt.Sprintf("mysql://%s:%s@%s:%s/%s", user, pass, host, port, db)
	case "3": // SQLite
		dsn = fmt.Sprintf("sqlite3://%s", db)
	case "4": // Other
		dsn = fmt.Sprintf("%s:%s@%s:%s/%s?sslmode=%s", user, pass, host, port, db, ssl)
	default:
		panic("Invalid option.")
	}

	// Print the DSN
	fmt.Println()
	white.Println("DSN:")
	yellow.Println(dsn)
	fmt.Println()
}
