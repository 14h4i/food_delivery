package common

import "fmt"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

func Recovery() {
	if err := recover(); err != nil {
		fmt.Println("Recovered from ", err)
	}
}
