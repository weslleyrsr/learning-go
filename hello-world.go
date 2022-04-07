package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	// "reflect"
)

var logs = []string{}

const checks = 10

const delay = 2

func main() {
	welcome()

	for {
		showOptions()
		handleOption(readOption())
	}
}

func welcome() {
	name := "Weslley"
	version := 1.1

	fmt.Println("Hi", name)
	fmt.Println("This program is at version:", version)
}

func showOptions() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Display logs")
	fmt.Println("0 - Exit program")
}

func readOption() int {
	option := 0

	fmt.Scan(&option)
	fmt.Println("Choosen option:", option)

	return option
}

func handleOption(option int) {
	switch option {
	case 0:
		exitProgram()
	case 1:
		observeLinks()
	case 2:
		displayLogs()
	default:
		restartProgram()
	}
}

func exitProgram() {
	fmt.Println("Exiting program...")
	os.Exit(0)
}

func observeLinks() {
	sites := readLinksFromFile()

	saveLog(
		"\n" + strings.Join(sites, "\n") + "\n",
	)

	for i := 0; i < checks; i++ {

		index := i + 1

		fmt.Println("Start monitoring round", index)

		saveLog(
			fmt.Sprint("\n", index, " - ", time.Now().Format("02/01/2006 15:04:05")),
		)

		for _, site := range sites {
			response, error := http.Get(site)

			if error == nil {
				message := fmt.Sprint("Try ", index, " got code ", response.Status)

				saveLog(
					fmt.Sprint("Online | ", response.Status),
				)

				fmt.Println(message)
			} else {
				saveLog(
					fmt.Sprint("Offline | ", response.Status),
				)

				fmt.Println("An error occured", error)
			}
		}

		time.Sleep(delay * time.Second)
	}

	fmt.Println("Monitoring complete, check logs for more details")
}

func displayLogs() {
	file, error := ioutil.ReadFile("log.txt")

	if error != nil {
		fmt.Println("An error occured", error)
	} else {
		fmt.Printf("File contents:\n%s", file)
	}
}

func restartProgram() {
	fmt.Println("Command not recognized!")
	// os.Exit(-1)
}

func readLinksFromFile() []string {
	var sites []string

	file, error := os.Open("links.txt")

	if error != nil {
		fmt.Println("An error occured", error)
		return sites
	}

	reader := bufio.NewReader(file)

	for {
		row, error := reader.ReadString('\n')
		sites = append(sites, strings.TrimSpace(row))

		if error != nil {
			break
		}
	}

	file.Close()

	return sites
}

func saveLog(log string) {
	file, error := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if error != nil {
		fmt.Println("An error occured", error)
	} else {
		file.WriteString(log + "\n")
		file.Close()
	}
}
