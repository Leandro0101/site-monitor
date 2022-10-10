package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func PutLog(site string, status bool){
	file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	CatchError(err)

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func PrintLogs(){
	file, err := ioutil.ReadFile("log.txt")
	CatchError(err)
	fmt.Println(string(file))
}