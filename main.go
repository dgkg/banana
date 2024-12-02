package main

import (
	"fmt"
	"os"
	"time"

	ser "github.com/dgkg/banana/services"
)

const StatusCodeOk = 200

const (
	StatusCodeCreated restStatus = iota + 201
	StatusCodeNoContent
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
	ReadFile()
	// execCode()
}

func execCode() error {
	fmt.Println(ser.Add(1, 2))

	var u myString = "user"
	fmt.Printf("%T - %v\n", u, u)

	res, err := ser.Sub("1", "2")
	if err != nil {
		return err
	}
	fmt.Println(res)
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	if res := ser.Add(1, 2); res != 3 {
		fmt.Println(err)
		err = generateError()
		if err != nil {
			return err
		}
	}

	fmt.Printf("%T - %v\n", StatusCodeOk, StatusCodeOk)
	fmt.Printf("%T - %v\n", StatusCodeCreated, StatusCodeCreated)
	fmt.Printf("%T - %v\n", StatusCodeNoContent, StatusCodeNoContent)

	var s string
	fmt.Printf("%T - `%v`\n", s, s)
	var i int
	fmt.Printf("%T - %v\n", i, i)

	m := make(map[string]int)
	m["toto"] = 1
	val, ok := m["toto"]
	fmt.Printf("%T - %v - %v\n", m, val, ok)
	delete(m, "toto")
	val, ok = m["toto"]
	fmt.Printf("%T - %v - %v\n", m, val, ok)

	mCal.Add(1, 2)
	mCal.Sub("1", "2")

	var usr User
	usr.id = 1
	usr.Name = "toto"
	usr.BirthDate = time.Now()
	usr.CreatedAt = time.Now()
	usr.UpdatedAt = time.Now()

	return nil
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

func ReadFile() error {
	defer fmt.Println("defer")
	data, err := os.ReadFile("toto.json")
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	fmt.Println(string(data))
	return nil
}

type myCalculator = ser.Calculator

var mCal myCalculator

type User struct {
	id   int
	Name string
	Human
	BddFields
}

type Human struct {
	BirthDate time.Time
}

type BddFields struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}
