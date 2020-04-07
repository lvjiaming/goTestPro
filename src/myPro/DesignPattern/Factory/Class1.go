// 第一个类
package Factory

import "fmt"

func init()  {
	Register("class1", func() Class {
		return new(Class1)
	})
}

type Class1 struct {

}

func (c *Class1) Do ()  {
	fmt.Println("Class1")
}


