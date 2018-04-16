package main

import (
	"fmt"
	"reflect"
)

type Service interface {
	fun1()
}

type ServiceImpl struct {
	Status string `bean:"carDao"`
}

func (x ServiceImpl) Fun1() {
	fmt.Println("i am in x.fun1")
}

type TypeRegister map[string]reflect.Type

func (t TypeRegister) Set(name string, value reflect.Type) {
	//name string, typ reflect.Type
	t[name] = value
}

func (t TypeRegister) Get(name string) reflect.Type {
	if v, ok := t[name]; ok {
		return v
	}
	return nil
}

var typeReg = make(TypeRegister)

func init() {
	s := ServiceImpl{}
	typeReg.Set("myservice", reflect.TypeOf(s))
}

func Destroy(subj interface{}) {
	stype := reflect.ValueOf(subj).Elem()
	fmt.Printf("stype ==== %T: %#v\n", stype, stype)
	//field := stype.FieldByName("Status")
	//if field.IsValid() {
	//	field.SetString("9999999999")
	//}
	tt := reflect.TypeOf(subj).Elem()
	for i := 0; i < tt.NumField(); i++ {
		ff := tt.Field(i)
		fmt.Println(ff.Tag.Get("bean"));
		stype.Field(i).SetString("88888888")
	}
	// fmt.Printf("field: %+v\n", field.CanSet())
}

type action func(string) (string)

func main() {
	var a action = func(str string) string {
		fmt.Println("fun call")
		return "mystr"
	}
	a("fff")

	//var i interface{}
	//i := reflect.New(typeReg.Get("myservice"))
	var i interface{} = &ServiceImpl{}
	fmt.Printf("%T:%#v\n", i, i)
	Destroy(i)
	tt := typeReg.Get("myservice")
	//fmt.Printf("t ==== %T: %#v\n", t, t)
	//tt := reflect.TypeOf(t)
	j := reflect.New(tt)
	fmt.Printf("j ==== %T: %#v\n", j, j)
	//stype := reflect.TypeOf(&j)
	//fmt.Printf("stype ==== %T: %#v\n", stype, stype)
	for i := 0; i < tt.NumField(); i++ {
		ff := tt.Field(i)
		fmt.Printf("sss==%d %s\n", i, ff.Tag.Get("bean"))
	}
	//stype.Field(i).SetString("88888888")
	field := j.Elem().FieldByName("Status")
	field.SetString("88888")
	//Destroy(j)
	fmt.Printf("%T:%#v\n", j, j)

	//method := reflect.ValueOf(j).MethodByName("Fun1")
	method := j.MethodByName("Fun1")
	method.Call([]reflect.Value{})
}
