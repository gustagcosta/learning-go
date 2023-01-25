package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Welcome to Hello World with Golang")

	names := []string{}

	for {

		showOptions()

		input := getInput()

		switch input {
		case 1:
			printNames(names)
		case 2:
			names = addName(names)
		case 3:
			searchNames(names)
		case 9:
			fmt.Println("leaving...")
			os.Exit(0)
		default:
			fmt.Println("an error has occurred...")
			os.Exit(-1)
		}
	}
}

func printNames(names []string) {
	if len(names) == 0 {
		fmt.Println("no names founded")
		return
	}

	for i, name := range names {
		fmt.Println("============")
		fmt.Println(i+1, "-", name)
		fmt.Println("============")
	}
}

func showOptions() {
	fmt.Println("1 _ See names")
	fmt.Println("2 _ Add new name")
	fmt.Println("3 _ Search names on Github")
	fmt.Println("9 _ Leave")
}

func getInput() int {
	var input int

	fmt.Scanln(&input)

	return input
}

func addName(names []string) []string {
	fmt.Print("Type the new name: ")
	name := ""
	fmt.Scan(&name)
	names = append(names, name)
	return names
}

func searchNames(names []string) {
	fmt.Println("============")

	if len(names) == 0 {
		fmt.Println("no names founded")
		return
	}

	for _, name := range names {
		url := "https://api.github.com/users/" + name
		response, error := http.Get(url)

		if error != nil {
			log.Fatal("Something went wrong")
		}

		if response.StatusCode == 200 {
			fmt.Println("User", name, "founded")
		}

		if response.StatusCode == 404 {
			fmt.Println("User", name, "not founded")
		}
	}

	fmt.Println("============")
}
