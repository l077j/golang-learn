package main

import "fmt"

// 1.检查i是否为nil
// t := i(T)  这个表达式可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，就会返回值给 t，如果断言失败，就会触发 panic
// func main() {
// 	var i interface{} = 10
// 	t1 := i.(int)
// 	fmt.Println(t1)
// 	fmt.Println("===================分割线=====================")
// 	t2 := i.(string)
// 	fmt.Println(t2) // panic
// 	// 如果要断言的接口值是 nil，也会触发panic
// }

// 2.检查 i 存储的值是否为某个类型
// t, ok:= i.(T)  这个表达式也是可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，就会返回其值给 t，并且此时 ok 的值 为 true，表示断言成功
// 如果接口值的类型，并不是我们所断言的 T，就会断言失败，但和第一种表达式不同的事，这个不会触发 panic，而是将 ok 的值设为 false ，表示断言失败，此时t 为 T 的零值。
// func main() {
// 	var i interface{} = 10
// 	t1, ok := i.(int)
// 	fmt.Printf("%d-%t\n", t1, ok)
// 	fmt.Println("==============1==============")
// 	t2, ok := i.(string)
// 	fmt.Printf("%s-%t\n", t2, ok)
// 	fmt.Println("==============2==============")
// 	var k interface{}
// 	t3, ok := k.(interface{})
// 	fmt.Println(t3, "-", ok)
// 	fmt.Println("==============3==============")
// 	k = 10
// 	t4, ok := k.(interface{})
// 	fmt.Printf("%d-%t\n", t4, ok)
// 	t5, ok := k.(int)
// 	fmt.Printf("%d-%t\n", t5, ok)
// }

// type switch断言（区分多种类型）
func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}

func main() {
	findType(10)
	findType("hello")

	var k interface{}
	findType(k)

	findType(10.23)
}
