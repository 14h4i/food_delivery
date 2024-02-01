package common

import "fmt"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

const CurrentUser = "user"

func Recovery() {
	if err := recover(); err != nil {
		fmt.Println("Recovered from ", err)
	}
}

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
