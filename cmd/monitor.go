package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"site-monitor/pkg"
	"strings"
	"time"
)

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
	pkg.CatchError(err)
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso")
		pkg.PutLog(site, true)
	} else {
		fmt.Println("Site: ", site, "está com problemas, ", "status code: ", resp.StatusCode)
		pkg.PutLog(site, false)
	}
}

func readSites() []string {
	var sites []string 
	file, err := os.Open("sites.txt")
	pkg.CatchError(err)
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