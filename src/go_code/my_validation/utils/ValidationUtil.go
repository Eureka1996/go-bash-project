package utils

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

/**
 * @author wufuqiang
 * @date 2022/11/7 19:32
 */

type User struct {
	Name  string
	Age   int
	Email string
}

//凡是实现了Validatable接口的类型都是可校验的类型。
//validation.Validate()函数在校验某个类型的数据时，先校验传入该函数的所有规则。
//如果这些规则都通过了，那么Validate()函数判断该类型有没有实现Validatbale接口。
//如果实现了，则调用其Validate()方法进行校验。

func (u *User) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(2, 10)),
		validation.Field(&u.Age, validation.Required, validation.Min(1), validation.Max(200)),
		validation.Field(&u.Email, validation.Required, validation.Length(10, 50), is.Email))
	return err
}

// 使用函数Validate()对基本类型值进行校验
func TestValidate() {
	name := "wufuqiang"
	err := validation.Validate(name,
		validation.Required,
		validation.Length(2, 10),
		is.URL)
	fmt.Println(err)
}

// ValidateStruct()接受一个结构体的指针作为第一个参数，后面依次指定各个字段的规则。
func TestValidateStruct() {
	u1 := &User{
		Name: "wufuqiang",
		Age:  18,
		//Email: "wufuqiang@qq.com",
		Email: "wufuqiang's email",
	}
	fmt.Println("user1:", ValidateUser(u1))
}

func ValidateUser(u *User) error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(2, 10)),
		validation.Field(&u.Age, validation.Required, validation.Min(1), validation.Max(200)),
		validation.Field(&u.Email, validation.Required, validation.Length(10, 50), is.Email))
	return err
}

func ValidateMapUser(u map[string]interface{}) error {
	err := validation.Validate(u, validation.Map(
		validation.Key("name", validation.Required, validation.Length(1, 10)),
		validation.Key("age", validation.Required, validation.Min(1), validation.Max(200)),
		validation.Key("email", validation.Required, validation.Length(10, 50), is.Email),
	))
	return err
}

func TestSelfValidate() {

	u1 := &User{
		Name:  "wufuqiang",
		Age:   70,
		Email: "wufuqiang@qq.com",
	}
	fmt.Println("user1:", validation.Validate(u1))

	u2 := &User{
		Name:  "wufuqiang",
		Age:   70,
		Email: "wufuqiang's email",
	}
	fmt.Println("user2:", validation.Validate(u2))
}
