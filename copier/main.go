// copier 是一个用于在go中进行类型复制的库
// 特点：深拷贝， 复制同名的字段， 复制切片， 复制map， 复制方法
// copier的复制依赖于反射，所以性能上会有一定的损失，还有一类基于代码生成、通过生成类型转换的代码的复制库，这种方法不会损耗性能（jmattheis/goverter）

// copier 只对外暴露两个函数：copier.Copy 和 copier.CopyWithOption

package main

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Id   string
	Name string
	// 当作为目标结构体时，忽略该字段
	Address string `copier:"-"`
}

type Student struct {
	// 指定字段名
	StudentId   string `copier:"Id"`
	StudentName string `copier:"Name"`
	Address     string
	School      string
	Class       string
}

// 1.不同类型结构体转换
// func main() {
// 	student := Student{
// 		StudentId:   "123",
// 		StudentName: "jack",
// 		Address:     "usa",
// 		School:      "MIT",
// 		Class:       "AI",
// 	}
// 	user := User{}
// 	if err := copier.Copy(&user, &student); err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", student)
// 	fmt.Printf("%+v\n", user)
// }

// 2.复制切片
// func main() {
// 	student := []Student{
// 		{
// 			StudentId:   "123",
// 			StudentName: "jack",
// 			Address:     "usa",
// 			School:      "MIT",
// 			Class:       "AI",
// 		},
// 		{
// 			StudentId:   "123",
// 			StudentName: "jack",
// 			Address:     "usa",
// 			School:      "MIT",
// 			Class:       "AI",
// 		},
// 	}

// 	var user []User
// 	if err := copier.Copy(&user, &student); err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", student)
// 	fmt.Printf("%+v\n", user)
// }

// 3.复制map
// func main() {
// 	student := Student{
// 		StudentId:   "123",
// 		StudentName: "jack",
// 		Address:     "usa",
// 		School:      "MIT",
// 		Class:       "AI",
// 	}

// 	src := make(map[string]Student)
// 	src["a"] = student
// 	src["b"] = student

// 	dest := make(map[string]User)

// 	if err := copier.Copy(&dest, &src); err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", src)
// 	fmt.Printf("%+v\n", dest)
// }

// 4.自定义转换方法
func main() {
	student := Student{
		StudentId:   "123",
		StudentName: "jack",
		Address:     "usa",
		School:      "MIT",
		Class:       "AI",
	}

	src := make(map[string]Student)
	src["a"] = student
	src["b"] = student

	dest := make(map[string]User)

	/*
		copier.CopyWithOption 是 copier 模块中更高级的克隆函数，它允许您使用在执行克隆操作时指定的一组选项对源和目标变量进行更精细的控制。以下是 copier.CopyWithOption 函数的一些选项：
		DeepCopy 选项：这个选项支持嵌套结构的深度复制。如果使用 DeepCopy 选项，则会递归复制嵌套结构的所有字段。
		IgnoreEmpty 选项：如果目标字段为空，则此选项将跳过源字段，而不是将其设置为空值。
		CopySymbolicLink：如果目标包含符号链接，则此选项将复制链接，而不是目标实际引用的内容。
		MatchUnexported：匹配未导出的结构体字段。
		TranslateFunc：这个选项允许用户指定一个转换函数，以在源和目标变量之间进行更高级的转换。例如，您可以使用该选项将时间字符串转换为 time.Time 类型。
		IgnoreFields：此选项允许您指定要在源或目标结构体中排除的字段。这个选项可以用于在目标结构体中创建部分复制，而不是复制整个结构体。
		AddPermission：将特定的权限添加到目标变量中。
		ReplaceSlice：当目标字段未初始化时替换切片。
		IgnoreCase：在克隆时忽略字段名称的大小写。
	*/
	if err := copier.CopyWithOption(&dest, &src, copier.Option{
		IgnoreEmpty:   false,
		CaseSensitive: false,
		DeepCopy:      false,
		// Converters: [...] 是一个对自定义转换器函数的定义。
		// 这里定义了一个新的 TypeConverter 对象并将其添加到 Converters 数组中
		Converters: []copier.TypeConverter{
			{
				SrcType: Student{}, // 定义了源类型 Student。我们的转换器函数使用此类型作为源值的输入参数类型。
				DstType: User{},    // 定义了目标类型 User。我们的转换器函数使用此类型作为目标值的输出参数类型。
				// Fn: func(...){...} 定义了自定义转换器函数的具体实现。该函数用于将源结构体的值转换为目标结构体的值
				// 转换时需要将源值强制转换为自定义类型 Student，以便从中获取所需的字段信息。
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(Student)
					if !ok {
						return User{}, errors.New("error type")
					}
					return User{
						Id: s.StudentId,
					}, nil
				},
			},
		},
		// FieldNameMapping: nil 是一个字典类型参数，用于将源字段与对应的目标字段进行映射。
		// 设置为 nil，表示源和目标结构体的字段名保持相同
		FieldNameMapping: nil,
	}); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", src)
	fmt.Printf("%+v\n", dest)
}
