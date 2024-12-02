package main

import (
	"fmt"

	ser "github.com/dgkg/banana/services"
)

const StatusCodeOk = 200

const (
	StatusCodeCreated   restStatus = 201
	StatusCodeNoContent restStatus = 204
)

type restStatus int

func (rs restStatus) String() string {
	switch rs {
	case StatusCodeCreated:
		return "Created"
	case StatusCodeNoContent:
		return "No Content"
	}
	return fmt.Sprintf("Status code: %d", rs)
}

func init() {
	initApp()
}

func initApp() {
	fmt.Println("init")
}

func main() {
	fmt.Println(ser.Add(1, 2))

	var u myString = "user"
	fmt.Printf("%T - %v\n", u, u)

	res, err := ser.Sub("1", "2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	if res := ser.Add(1, 2); res != 3 {
		fmt.Println(err)
		err = generateError()
	}

	fmt.Printf("%T - %v\n", StatusCodeOk, StatusCodeOk)
	fmt.Printf("%T - %v\n", StatusCodeCreated, StatusCodeCreated)
	fmt.Printf("%T - %v\n", StatusCodeNoContent, StatusCodeNoContent)
}

func generateError() error {
	return fmt.Errorf("error")
}

type myString string

func fibo() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
