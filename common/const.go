package common

import "fmt"

func Recovery() {
	if err := recover(); err != nil {
		fmt.Println("Recovered from ", err)
	}
}
