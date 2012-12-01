package main

import(
	"fmt"
	"reflect"
)

type My struct{
	name string
}

func (this *My) GetName() string{
	return this.name
}

func (this *My) SetName(s string){
	this.name = s
}

func main() {
	s := "This is string"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))	
	fmt.Println("------------------")

	var x float64 = 3.4
	fmt.Println(reflect.ValueOf(x))
	fmt.Println("------------------")

	a := new(My)
	a.name = "likun"
	t := reflect.TypeOf(a)
	fmt.Println(t.NumMethod())
	fmt.Println("------------------")

	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(b[0])

	//d := reflect.ValueOf(a).MethodByName("SetName").Call([]reflect.Value{})
	//fmt.Println(d[0]("LiKun"))
}
