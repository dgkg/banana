package pointers

import "fmt"

func Pointer() {
	var a int = 42
	var b *int = &a
	var c **int = &b

	//fmt.Printf("%T\n", *a)
	fmt.Printf("%T\n", c)

	fmt.Println(a, b)
	fmt.Println(&a, *b)
	fmt.Printf("%T\n", b)

	var u *User = &User{
		Name: "Lucas",
	}
	fmt.Printf("%T\n", u)
	u = nil
	//fmt.Printf("%T\n", *u)

	var list []interface{}
	list = append(list, "a", u, a, b, c)
	fmt.Println(list)
	fmt.Printf("%T\n", list)
}

type User struct {
	Name string
}
