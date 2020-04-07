package ExampleCollection

import (
	"fmt"
	"github.com/bwmarrin/snowflake" // 需要安装（go get -u github.com/bwmarrin/snowflake）
	"github.com/sony/sonyflake" // 安装同上
	"time"
)
/**
 id生成器（
 */

// Snowflakes实现
func CreateIdBySnowflake()  {
	n, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < 4; i++ {
		id := n.Generate()
		fmt.Println("id: ", id)
		fmt.Println("node: ", id.Node(), "step: ", id.Step(), "time: ", id.Time())
	}
}

// sonyflake实现
func CreateIdBySonyflake()  {
	t, _ := time.Parse("2006-01-02", "2018-01-01")
	settings := sonyflake.Settings{
		StartTime: t,
	}
	/*
	type Settings struct {
	    StartTime time.Time // 如果不设置的话，默认是从 2014-09-01 00:00:00 +0000 UTC 开始。
	    MachineID func() (uint16, error) 可以由用户自定义的函数，如果用户不定义的话，会默认将本机 IP 的低 16 位作为 machineid。
	    CheckMachineID func(uint16) bool 是由用户提供的检查 MachineID 是否冲突的函数
	}
	 */
	s := sonyflake.NewSonyflake(settings)
	for i := 0; i < 3; i++ {
		id, err := s.NextID()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("id: ", id)
	}
}

/**
 map的多键索引
 */
// 利用哈希值
type student struct {
	name string
	age int
	sex string
}
var mapper = make(map[int][]student)  // 映射表
type classicQueryKey struct {
	name string
	age int
}
func (c *classicQueryKey) hash() int {
	return simpleHash(c.name) + c.age * 1000000
}

func simpleHash(name string) (ret int) {
	for i := 0; i < len(name); i++ {
		c := name[i]
		ret += int(c)
	}
	return
}
func buildIndex(stu []student) map[int][]student { // 构建映射表
	for _, s := range stu{
		mapKey := classicQueryKey{name: s.name, age:  s.age,}
		mapVal := mapper[mapKey.hash()]

		mapVal = append(mapVal, s)

		mapper[mapKey.hash()] = mapVal
	}
	return mapper
}
func queryData(name string, age int) []student { // 查找
	mapKey := classicQueryKey{name: name, age: age}
	mapVal := mapper[mapKey.hash()]
	return mapVal
}
func CheckValByIndexUseHash()  {
	s := []student{
		{name: "张三", age: 10},
		{name: "李四", age: 21},
		{name: "王五", age: 20},
		{name: "王麻子", age: 10},
		{name: "王麻子", age: 10},
	}
	mapper := buildIndex(s)
	fmt.Println(mapper)
	checkVal := queryData("王麻子", 10)
	fmt.Println("查找结果：", checkVal)
}