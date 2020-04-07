// 工厂模式
package Factory

type Class interface {
	Do()
}

var factoryList map[string]func() Class = make(map[string]func() Class)

func Register(name string, item func() Class)  {
	factoryList[name] = item
}

func GetClassByName(name string) func() Class {
	item := factoryList[name]
	return item
}

func Create(name string) Class {
	if i := factoryList[name]; i != nil {
		return i()
	} else {
		panic("not find class")
	}
}