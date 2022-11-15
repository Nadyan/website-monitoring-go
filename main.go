package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
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
		printLogs()
	case 3:
		fmt.Println("\nExiting app...")
		os.Exit(0)
	default:
		fmt.Println("Option not available")
	}
}

func initMonitoring() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("\nMonitoring websites...")

	websites := readWebsitesFile()
	for _, site := range websites {
		res, err := http.Get(site)
		printAndSaveStatus(site, res, err)
	}
}

func readWebsitesFile() []string {
	var websites []string

	file, err := os.Open("files/websites.txt")
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

func printAndSaveStatus(site string, res *http.Response, err error) {
	status := "Online"

	if res == nil || res.StatusCode != 200 {
		status = "Offline"
	}

	fmt.Println(site, "->", status)
	saveLog(site, status)
}

func saveLog(site string, status string) {

	timestamp := time.Now().Local().String()

	file, err := os.OpenFile("files/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error saving in log: ", err)
		return
	}

	file.WriteString(timestamp + " " + site + " -> " + status + "\n")

	file.Close()
}

func printLogs() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("\nPrevious status:")

	file, err := ioutil.ReadFile("files/log.txt")
	if err != nil {
		fmt.Println("Error reading logs:", err)
		return
	}

	fmt.Println(string(file))
}
