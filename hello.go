package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
			printLogs()
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

func initMonitoring(){
	fmt.Println("Monitoring...")
	sites := readSites()
	fmt.Println(sites)

	for i := 0 ; i < 5; i++ {
		for _, site:= range sites {
			fmt.Println(site)
			testSite(site)
		}
		fmt.Println()
		time.Sleep(2 * time.Second)
	}
}

func testSite(site string) {
	resp, err := http.Get(site)
	catchError(err)
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso")
		putLog(site, true)
	} else {
		fmt.Println("Site: ", site, "está com problemas, ", "status code: ", resp.StatusCode)
		putLog(site, false)
	}
}

func readSites() []string {
	var sites []string 
	file, err := os.Open("sites.txt")
	catchError(err)
	// file, err := ioutil.ReadFile("sites.txt")
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if(err == io.EOF){
			break
		}

		sites = append(sites, line)
	}

	file.Close()
	
	return sites
}

func catchError(err error) {
	if(err != nil) {
		fmt.Println("Catched error: ", err)
	}
}

func putLog(site string, status bool){
	file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	catchError(err)

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs(){
	file, err := ioutil.ReadFile("log.txt")
	catchError(err)
	fmt.Println(string(file))
}