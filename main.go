package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var option int
	for option != 3 {
		option = get_option_menu()
		executeOption(option)
	}
}

func get_option_menu() int {
	fmt.Println("\n1 - Init website scan")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")

	var optionInput int
	fmt.Scan(&optionInput)

	return optionInput
}

func executeOption(option int) {
	switch option {
	case 1:
		initMonitoring()
	case 2:
		fmt.Println("Executing 2")
	case 3:
		fmt.Println("\nExiting app...")
		os.Exit(0)
	default:
		fmt.Println("Option not available")
	}
}

func initMonitoring() {
	fmt.Println("\nMonitoring websites...")

	websites := readWebsitesFile()
	for _, site := range websites {
		res, err := http.Get(site)
		printResult(site, res, err)
	}
}

func readWebsitesFile() []string {
	var websites []string

	file, err := os.Open("websites.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return websites
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
		}
		websites = append(websites, strings.TrimSpace(line))
	}

	file.Close()

	return websites
}

func printResult(site string, res *http.Response, err error) {
	greenColor := "\033[32m"
	redColor := "\033[31m"
	resetColor := "\033[0m"

	status := greenColor + "Online" + resetColor

	if res == nil || res.StatusCode != 200 {
		status = redColor + "Offline" + resetColor
	}

	fmt.Println(site, "->", status)

}
