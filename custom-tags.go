package main

import "fmt"
import "reflect"

type User struct {
	Name  string `mytag:"MyName"`
	Email string `mytag:"MyEmail"`
}

func main() {
	u := User{"Bob", "bob@example.com"}
	t := reflect.TypeOf(u)
	t1 := reflect.TypeOf(&u).Elem()
	t2 := reflect.TypeOf(u)//.Elem()
	v1 := reflect.ValueOf(&u).Elem()
	v2 := reflect.ValueOf(u).Interface()

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(v1)
	fmt.Println(v2)

	if t1 == t2{
		fmt.Println("t1 == t2")
	}
	if v1 == v2{
		fmt.Println("v1 == v2")
	}
	for _, fieldName := range []string{"Name", "Email"} {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}

		reflect.ValueOf(&u).Elem().FieldByName(fieldName).Set(reflect.ValueOf("abc"))
		fmt.Printf("\nField: User.%s\n", fieldName)
		fmt.Printf("\tWhole tag value : %q\n", field.Tag)
		fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))
	}
}
