package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/dgkg/banana/model"
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
	// fmt.Println("init")
}

func main() {

	// wg := &sync.WaitGroup{}
	// emailNum := 10000
	// wg.Add(emailNum)
	// for i := 0; i < emailNum; i++ {
	// 	go SendEmail(wg)
	// }
	// wg.Wait()

	u := model.CreateUserWithPtr("Doe")
	fmt.Printf("type: %T - value: %#v - pointer: %p\n\n", u, u, u)
	u2 := model.CreateUserByValue("Bob")
	fmt.Printf("type: %T - value: %#v - pointer: %p\n\n", u2, u2, u2)
	// ReadFile()
	// execCode()

	// for i := 0; i < 100; i++ {
	// 	var j uint8
	// 	fmt.Printf("%p\n", &j)
	// }

	tblMatrix := [3][3][3]int{
		{{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}},
		{{10, 11, 12},
			{13, 14, 15},
			{16, 17, 18}},
		{{10, 11, 12},
			{13, 14, 15},
			{16, 17, 18}},
	}

	fmt.Println(tblMatrix)

	tblA := []int{1, 2, 3, 4, 5}
	fmt.Println(tblA)
	// tblB := make([]int, len(tblA))
	// copy(tblB, tblA)
	var tblB []int
	var tblC []int
	tblB = append(tblB, tblA...)
	tblC = append(tblC, tblA[:]...)

	fmt.Printf("%p - %v - %v - %v\n", &tblA, tblA, len(tblA), cap(tblA))
	fmt.Printf("%p - %v - %v - %v\n", &tblB, tblB, len(tblB), cap(tblB))
	fmt.Printf("%p - %v - %v - %v\n", &tblC, tblC, len(tblC), cap(tblC))

	tblA[0] = 10
	fmt.Printf("%p - %v - %v - %v\n", &tblA, tblA, len(tblA), cap(tblA))
	fmt.Printf("%p - %v - %v - %v\n", &tblB, tblB, len(tblB), cap(tblB))
	fmt.Printf("%p - %v - %v - %v\n", &tblC, tblC, len(tblC), cap(tblC))
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

func SendEmail(wg *sync.WaitGroup) error {
	time.Sleep(5 * time.Second)
	fmt.Println("Email sent")
	wg.Done()
	return nil
}

func ReadFile() error {
	defer fmt.Println("defer")
	// data, err := ioutil.ReadFile("toto.json")
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
