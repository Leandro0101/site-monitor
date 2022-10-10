package pkg

import "fmt"

func CatchError(err error) {
	if(err != nil) {
		fmt.Println("Catched error: ", err)
	}
}

