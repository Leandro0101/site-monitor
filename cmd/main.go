package main

import (
	"fmt"
	"os"
	"site-monitor/pkg"
)

func main(){
	name := "Leandro"
	version := 1.0
	fmt.Println("Hello, ", name)
	fmt.Println("The program version is", version)


	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			initMonitoring()
		case 2: 
			fmt.Println("Showing logs...")
			pkg.PrintLogs()
		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default :
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}

func readCommand () int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen command was ", command)
	return command
}

func showMenu() {
	fmt.Println("1 - Init monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit program")
}


