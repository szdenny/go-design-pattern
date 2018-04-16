package main

import "fmt"
import "reflect"

type T struct {}

func (t *T) Foo(){
    fmt.Println("foo")
}

func main(){
    t := T{}
    reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})
}
