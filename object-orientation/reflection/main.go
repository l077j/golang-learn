package main

import (
	"fmt"
	"reflect"
)

// func main() {
// 	var age interface{} = 25
// 	fmt.Printf("原始接口变量的类型为%T,值为%v\n", age, age)
// 	t := reflect.TypeOf(age)
// 	v := reflect.ValueOf(age)

// 	// 从接口变量到反射对象
// 	fmt.Printf("从接口变量到反射对象:Type对象的类型为 %T \n", t)
// 	fmt.Printf("从接口变量到反射对象:Value对象的类型为 %T \n", v)

// 	// 从反射对象到接口变量
// 	i := v.Interface()
// 	fmt.Printf("从反射对象到接口变量:新对象的类型为 %T 值为 %v \n", i, i)

// 	// 不是接收变量指针创建的反射对象，是不具备『可写性』的
// 	var name string = "matt"
// 	v1 := reflect.ValueOf(name)
// 	fmt.Println("可写性为:", v1.CanSet())

// 	// 要让反射对象具备可写性，需要注意1.创建反射对象时传入变量的指针2.使用 Elem()函数返回指针指向的数据
// 	var name2 string = "matt2"
// 	v2 := reflect.ValueOf(&name2)
// 	fmt.Println("v2 可写性为:", v2.CanSet())
// 	v3 := v2.Elem()
// 	fmt.Println("v3 可写性为:", v3.CanSet())

// 	// 可以对具有可写性的反射对象进行修改操作
// 	fmt.Println("真实世界里 name 的原始值为：", name2)
// 	v3.SetString("lj")
// 	fmt.Println("通过反射对象进行更新后，真实世界里 name2 变为：", name2)
// }

// type Profile struct {
// 	name   string
// 	age    int
// 	gender string
// }
// func main() {
// 	m := Profile{}

// 	t := reflect.TypeOf(&m)
// 	fmt.Println("Type:", t)
// 	fmt.Println("Kind:", t.Kind())

// 	fmt.Println("m Type:", t.Elem())
// 	fmt.Println("m Kind:", t.Elem().Kind())

// 	var age int = 25
// 	v1 := reflect.ValueOf(age)
// 	fmt.Printf("转换前, type: %T, value: %v\n", v1, v1)
// 	v2 := v1.Int()
// 	fmt.Printf("转换后, type: %T, value: %v\n", v2, v2)

// 	var score float64 = 99.5
// 	v3 := reflect.ValueOf(score)
// 	fmt.Printf("转换前, type: %T, value: %v\n", v3, v3)
// 	v4 := v3.Float()
// 	fmt.Printf("转换后, type: %T, value: %v\n", v4, v4)

// 	var name string = "matt"
// 	v5 := reflect.ValueOf(name)
// 	fmt.Printf("转换前, type: %T, value: %v\n", v5, v5)
// 	v6 := v5.String()
// 	fmt.Printf("转换后, type: %T, value: %v\n", v6, v6)

// 	var isMale bool = true
// 	v7 := reflect.ValueOf(isMale)
// 	fmt.Printf("转换前, type: %T, value: %v\n", v7, v7)
// 	v8 := v7.Bool()
// 	fmt.Printf("转换后, type: %T, value: %v\n", v8, v8)

// 	v9 := reflect.ValueOf(&age)
// 	fmt.Printf("转换前, type: %T, value: %v\n", v9, v9)
// 	v10 := v9.Pointer()
// 	fmt.Printf("转换后, type: %T, value: %v\n", v10, v10)

// 	v11 := reflect.ValueOf(name)
// 	fmt.Printf("转换前, type: %T, value: %v\n", v11, v11)
// 	v12 := v11.Interface()
// 	fmt.Printf("转换后, type: %T, value: %v\n", v12, v12)
// }

// func main() {
// 	var numList []int = []int{1, 2}

// 	v1 := reflect.ValueOf(numList)
// 	fmt.Printf("转换前, type: %T, value: %v\n", v1, v1)
// 	// Slice函数接收两个参数
// 	v2 := v1.Slice(0, 2)
// 	fmt.Printf("转换后, type: %T, value: %v\n", v2, v2)
// }

// func appendToSlice(arrPtr interface{}) {
// 	valuePtr := reflect.ValueOf(arrPtr)
// 	value := valuePtr.Elem()

// 	// 更新切片
// 	value.Set(reflect.Append(value, reflect.ValueOf(3)))

// 	fmt.Println(value)
// 	fmt.Println(value.Len())
// }
// func main() {
// 	arr := []int{1, 2}
// 	appendToSlice(&arr)
// 	fmt.Println(arr)
// }

// type Person struct {
// 	name   string
// 	age    int
// 	gender string
// }
// func (p Person) SayBye() {
// 	fmt.Println("Bye")
// }
// func (p Person) SayHello() {
// 	fmt.Println("Hello")
// }
// func main() {
// 	p1 := Person{"matt", 23, "male"}
// 	v := reflect.ValueOf(p1)
// 	fmt.Println("字段数:", v.NumField())
// 	fmt.Println("第一个字段:", v.Field(0))
// 	fmt.Println("第二个字段:", v.Field(1))
// 	fmt.Println("第三个字段:", v.Field(2))
// 	fmt.Println("================================================")
// 	fmt.Println(v)
// 	// 也可以这样来遍历
// 	for i := 0; i < v.NumField(); i++ {
// 		fmt.Printf("第%d个字段: %v \n", i+1, v.Field(i))
// 	}

// 	p2 := &Person{"lj", 22, "female"}
// 	t := reflect.TypeOf(p2)
// 	fmt.Println("方法数（可导出的）:", t.NumMethod())
// 	fmt.Println("第一个方法:", t.Method(0).Name)
// 	fmt.Println("第二个方法:", t.Method(1).Name)
// 	fmt.Println("===============================================")
// 	// 也可以这样来遍历
// 	for i := 0; i < t.NumMethod(); i++ {
// 		fmt.Printf("第%d个方法: %v \n", i+1, t.Method(i))
// 		fmt.Printf("第%d个方法: %v \n", i+1, t.Method(i).Name)
// 	}
// }

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) SayBye() string {
// 	return "Bye"
// }
// func (p Person) SayHello() string {
// 	return "Hello"
// }
// func main() {
// 	p := &Person{"matt-lj", 26}
// 	t := reflect.TypeOf(p)
// 	v := reflect.ValueOf(p)
// 	for i := 0; i < v.NumMethod(); i++ {
// 		fmt.Printf("调用第%d个方法:%v, 调用结果: %v\n", i+1, t.Method(i).Name, v.Elem().Method(i).Call(nil))
// 	}
// }

// type Person struct {
// 	name   string
// 	age    int
// 	gender string
// }

// func (p Person) SayBye() {
// 	fmt.Print("Bye")
// }
// func (p Person) SayHello() {
// 	fmt.Println("Hello")
// }
// func main() {
// 	p := &Person{"matt", 27, "male"}

// 	v := reflect.ValueOf(p)

// 	v.MethodByName("SayHello").Call(nil)
// 	v.MethodByName("SayBye").Call(nil)
// }

type Person struct {
}

func (p Person) SelfIntroduction(name string, age int) {
	fmt.Printf("Hello, my name is %s and i'm %d years old.", name, age)
}
func main() {
	p := &Person{}
	v := reflect.ValueOf(p)
	name := reflect.ValueOf("matt")
	age := reflect.ValueOf(23)
	input := []reflect.Value{name, age}
	v.MethodByName("SelfIntroduction").Call(input)
}
