package main // 所属包（main包且含有main为命令源码文件）

import (
	"archive/zip"
	"bytes"
	"container/list"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"io"
	"log"
	"math"
	"math/rand"
	"myPro/BinaryTree"
	"myPro/CustomerSystem/View"
	"myPro/DesignPattern/Factory"
	"myPro/ExampleCollection"
	"myPro/FiniteStateMachine"
	"myPro/MyDictionary"
	"myPro/SimpleMediaPlayer"
	"myPro/TestWebServer"
	"myPro/customLog"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main()  {
	//createPic()
	//pointInfo()
	//sliceInfo()
	//mapInfo()
	//listInfo()
	//sentenceInfo()
	//binaryFindTest()
	//chainTest()
	//interfaceInfo()
	//jsonTestFunc()
	//testListFunc()
	//testLogFunc()
	//testSortFunc()
	//DictionaryFunc()
	//typeSwitchTestFunc()
	//webServerFunc()
	//musicPlayerFunc()
	//testReflect()
	//FiniteStateMachineFunc()
	//BinaryTreeTest()
	//DesignPatternFun()
	//createGifAni()
	//timePackageFunc()
	//packageFunc()
	//customerSystem()
	//goroutineFunc()
	//goroutineTest()
	//goroutineFunc2()
	//goroutineFunc3()
	//goroutineFunc4()
	//goroutineFunc5()
	//Telent.TelnetSer("192.168.0.11:2001")
	//goroutineFunc6()
	//goroutineFunc7()
	//goroutineFunc8()
	//goroutineFunc9()
	//ChatServer.ChatServer()
	//jsonFileFunc()
	//xmlFileFunc()
	//gobFileFunc()
	//EncodeAndDecode.EncodeAndDecodeFunc()
	//TestWebServer.TWSHandlerFunc()
	//ExampleCollection.CreateIdBySonyflake()
	ExampleCollection.CheckValByIndexUseHash()
}


func createPic()  {
	size := 300
	pic := image.NewGray(image.Rect(0, 0, size, size))  // 创建一个图片数据
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{0})  // 将每个像素设为黑色
		}
	}
	file, err := os.Create("bg.png") // 创建一张图片文件
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, pic) //将image信息写入文件中
	// 关闭文件
	file.Close()
}
/**
 指针
 */
func pointInfo()  {
	tStr := "test_string"
	p := &tStr  // 获取地址
	v := *p // 获取此地址的值
	fmt.Printf("指针类型: %T\n", p)
	fmt.Printf("指针地址： %p\n", p)
	fmt.Printf("指针指向的值类型： %T\n", v)
	fmt.Printf("指针指向的值： %s\n", v)
	//创建指针
	newPoint := new(string)
	*newPoint = "10"
	value := *newPoint
	fmt.Printf("指针的值类型: %T\n", newPoint)
	fmt.Printf("指针的值： %s\n", value)
}

/**
 切片
 */
func sliceInfo()  {
	var a []int // 声明切片
	a = append(a, 1)
	a = append(a, 1, 2, 3)
	a = append(a, []int{1, 2, 3}...) // 切片中插入切片
	fmt.Println("切片：", a)
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("切片长度: %d  切片容量: %d 切片地址: %p\n", len(numbers), cap(numbers), numbers)
	}

	b := []int{1, 2, 3, 4, 5}
	c := 10
	// 将c插入b的第2个位置
	b = append(append(b[:2]), append([]int{c}, append(b[2:])...)...)
	fmt.Println(b)

	// 切片复制
	d := []int{1, 2, 3, 4, 5}
	e := []int{6, 7, 9}
	f := copy(d, e)
	fmt.Println(d, e, f)

	// 遍历切片（range返回两个值，一个为元素下标，一个为元素值）
	for _, value := range d{
		fmt.Println("切片值", value)
	}
}

/**
 map
 */
func mapInfo()  {
	var a map[int] string
	a = map[int]string{1: "10", 2: "20"}
	b := map[int]string{1: "1", 2: "2"}
	c := make(map[int]string)
	d := make(map[int][]int)
	e := []int{1, 3, 3}
	d = map[int][]int{1: e, 3: e, 2: []int{2, 3}}
	fmt.Println(a, b, c, d)
	delete(d, 2)
	fmt.Println("删除后的map:", d)

	// 并发map（在非并发情况下，前者性能佳）
	var f = sync.Map{}
	f.Store("hand", 2)
	f.Store("eys", 3)
	fmt.Println(f.Load("hand"))
	f.Delete("hand")
	f.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true // 返回true继续遍历，返回false终止遍历
	})
}

/**
 链表
 */
func listInfo()  {
	a := list.New()
	a.PushBack("10")
	a.PushFront(10)
	b := a.PushBack("newAdd")
	a.InsertAfter("这个之后", b)
	a.InsertBefore("这个之前", b)
	fmt.Println("删除之前")
	for i := a.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	a.Remove(b) // 只能通过添加进去返回的去删除
	fmt.Println("删除之后")
	for i := a.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

/**
 语句
 */
func test() bool {
	return true
}
func sentenceInfo()  {
	if a := test(); a { // 可加语句
		fmt.Println("条件成立")
	}
	b := 1
	for b < 5 {
		b++
		fmt.Println(b)
	}
	// 输出九九乘法表
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i * j)
		}
		fmt.Println()
	}
	c := 2
	switch  {
	case c > 1:
		fmt.Println("大于1")
		fallthrough  // 继续执行下列的
	case c == 2:
		fmt.Println("等于2")
	}
	switch c {
	case 2, 3:
		fmt.Println("两个分支")
	}

	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			if j == 5 {
				goto onExit  // 跳出循环，执行指定标签的操作
			}
		}
	}

	return
	onExit: // 定义的标签
		//fmt.Println("跳出循环了")
	for i := 0; i < 100; i++ {
		if i == 5 {
			break onExit // 跳出指定循环（标签）（continue 也可以如此做）
		}
	}
}

/**
 二分法查找
 */
func binaryFind(arr *[]int, min, max, value int)  {
	var findFunc func(index1, _ int)
	findFunc = func(start, end int) { // 匿名函数
		centerIndex := (start + end) / 2
		if (*arr)[centerIndex] == value {
			fmt.Println("找到了, 下表为：", centerIndex)
		} else {
			if (*arr)[centerIndex] > value {
				findFunc(min, centerIndex - 1)
			} else {
				findFunc(centerIndex + 1, max)
			}
		}
	}
	findFunc(min, max)
}

func binaryFindTest()  {
	arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92}
	binaryFind(&arr, 0, len(arr), 39)
}

/************************链式处理**************************/
func stringProccess(list *[]string, chain []func(string) string)  {
	for index, str := range *list{
		var result string
		for _, proc := range chain{
			result = proc(str)
		}
		(*list)[index] = result
	}
}
func chainTest()  {
	list := []string{
		"go Scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	chain := []func(string) string {  // 链式处理函数（函数切片）
		strings.TrimSpace,
		strings.ToUpper, // 转换为大写
	}
	fmt.Println("处理前：", list)
	stringProccess(&list, chain)
	fmt.Println("处理后：", list)
}
/************************链式处理**************************/

/***************************函数，接口，与结构体******************************/
type testStruct struct {  // 定义结构体

}

//利用结构体实现接口
type testInterface interface {  // 定义接口
	testOne(a string) string
	testTwo()
}

func (p *testStruct) testOne(a string) string {  // 前面是接收器，接受这个对象的实例()
	fmt.Println(a)
	return ""
}
func (p *testStruct) testTwo()  {
	fmt.Println("结构体实现接口的第二个函数")
}

// 利用函数实现接口
type testFunc func(p interface{})
type testInterfaceTwo interface {
	twoFunc(a interface{})
	twoFuncTwo()
}
func (p testFunc) twoFuncTwo () {
	fmt.Println("函数接口实现的第二个函数")
}
func (p testFunc) twoFunc (a interface{}) {
	fmt.Println(a)
	p(a)
}
//func (p testFunc) testTwo ()  {
//	fmt.Println("第二个接口函数")
//}

/**
 可变参数
 */
func myPrint(args ...interface{})  {
	for _, value := range args{ // 类型断言
		switch value.(type) {
		case int:
			fmt.Println(value, "is int value")
		case string:
			fmt.Println(value, "is string value")
		case int32:
			fmt.Println(value, "is int32 value")
		}
	}
}

/**
 延迟执行语句,函数执行完执行延迟语句（有多个时，逆顺序执行）
 */
func deferFunc()  {
	fmt.Println("语句1")
	defer fmt.Println("语句2")
	defer fmt.Println("语句3")
	defer fmt.Println("语句4")
	fmt.Println("语句5")
}

/**
 自定义error
 */
type customError struct {
	fileNam string
	fileLine int
}

func (p *customError) Error() string { // 实现error接口中的Error函数
	return fmt.Sprintf("%s:%d",p.fileNam, p.fileLine)
}
func testError(fileName string, fileLine int) error {// error的类型为结构体
	return &customError{fileName, fileLine}
}

/**
 宕机
 */
func panicFunc()  {
	fmt.Println("正常运行1")
	defer fmt.Println("延迟")
	panic("宕机")
	fmt.Println("正常运行2")
}
/**
 宕机与恢复宕机
 */
type panicFuncErrorInfo struct {
	info string
}

func panicDealWithFunc(f func())  {
	defer func() {
		err := recover() // 恢复宕机（会返回错误）
		switch err.(type) {
		case runtime.Error:
			fmt.Println("运行错误", err)
		default:
			fmt.Println("其他错误：", err)
		}
	}()
	f()
}

func testPanicRecover()  {
	fmt.Println("程序正常运行")
	panicDealWithFunc(func() {
		fmt.Println("手动宕机前")
		panic((&panicFuncErrorInfo{"手动引起宕机"}).info) // 通过取地址实例化panicFuncErrorInfo结构体
	})
	panicDealWithFunc(func() {
		fmt.Println("赋值前")
		var a *int
		*a = 1
		fmt.Println("赋值后")
	})
	fmt.Println("运行") // 宕机会跳出宕机处的函数
}
/**
程序运行时间（检测性能）
 */
func testXn()  {
	start := time.Now()
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	fmt.Println("用时：", time.Since(start))
}

/**
 type自定义类型，可以给任意类型添加方法
 */
type myInt int

func (p *myInt) IsZeo () bool { // 指针类型接收器
	return *p == 0
}
func (p myInt) Add (value int) int { // 非指针类型接收器
	return int(p) + value
}
func testTypeFunc(value myInt)  {
	var a myInt
	fmt.Println(a.IsZeo())
	a = value
	fmt.Println(a.IsZeo())
	fmt.Println(a.Add(10))
}
/**
 结构体
 */
type structOne struct {
	a, b int
}
type structTwo struct {
	structOne// 内嵌结构体，字段名就是类型名(内嵌结构体可用于实现继承关系)
	c int
}
type structThree struct {
	structOne
	a struct{  // 匿名内嵌结构体
		b int
		c string
	}
}

func testStructFunc()  {
	a := &structTwo{structOne{1, 2}, 3}
	fmt.Println(a)
	fmt.Println(a.a, a.b, a.c)
	fmt.Println(a.structOne.a, a.structOne.b, a.c)
	b := &structThree{ // 初始化内嵌结构体
		structOne: structOne{a: 1, b: 2},
		a: struct {
			b int
			c string
		}{b: 1, c: "111"},
	}
	fmt.Println(b)
	fmt.Println(b.a.b, b.structOne.a, b.structOne.b, b.a, b.a.c)
}
/**
接口
 */
type Server struct {
	Start func()
	Log func(interface{})
}
type LogServer struct {

}

func (l *LogServer) Log (s interface{})  {
	fmt.Println(s)
}

type GameServer struct { // 此结构体内嵌了LogServer结构体
	LogServer
}

func (g *GameServer) Start ()  {
	fmt.Println("开始游戏")
}
func interfaceTestFunc()  {
	a := new(GameServer)
	a.Log("哈哈")
	a.Start()
}
/**
 断言（i.(T): i为断言值，T为断言类型）
 */
func testDuanYan()  {
	var a interface{} = 10
	n, ok := a.(int) // 返回两个值。类似类型强转（第一个值未转换后的值，后一个值转换结果）
	fmt.Println(n, ok)
	m, result := a.(float64)
	fmt.Println(m, result)
	f, result := a.(string)
	fmt.Println(f, result)
	//f := a.(float64) // 此时断言未成功，且结果未呈现，会产生恐慌(即错误)
	//fmt.Println(f)
}

func interfaceInfo()  {
	type NewInterface interface {
		Test()  // 此时Test方法和接口名首字母都是大写，则此方法可以被接口所在的包（package）之外的代码访问。
		testTwo(a []string) (p string)
		testThree(string) string
	}
	a := new(testStruct)
	var b testInterface
	b = a// 此处的结构体已经实现了接口的所有方法，所以可以赋值
	b.testOne("测试呢")
	b.testTwo()
	var c testInterfaceTwo
	c = testFunc(func(p interface{}) {
		fmt.Println("函数本体", p)
	})
	c.twoFunc("111")
	c.twoFuncTwo()
	myPrint(1, "22")
	deferFunc()
	d := testError("main.go", 10)
	fmt.Println(d)
	//panicFunc()
	testPanicRecover()
	testXn()
	testTypeFunc(1)
	testStructFunc()
	interfaceTestFunc()
	testDuanYan()
	changeString("abcdef")
}


/**
 用于测试用例的函数
 */
func GetArea(weight int, height int) int {
	return weight * height
}
/***************************函数，接口，与结构体******************************/

/**
反射(对所有接口进行反射，都可以得到一个包含 Type 和 Value 的信息结构)
*/

type MyTestReflectType int
type MyTestReflectStruct struct {
	a int
	b string
}
type MyTestReflectUser struct {
	Name string
	Id int
	Sex int
}

func reflectTestCallFunc(a, b int) int {
	return a + b
}
/**
反射性能需注意如下
1.能使用原生代码时，尽量避免反射操作。
2.提前缓冲反射值对象，对性能有很大的帮助。
3.避免反射函数调用，实在需要调用时，先提前缓冲函数参数列表，并且尽量少地使用返回值。
 */
func testReflect()  {
	var a float64
	a = 3.14
	v := reflect.ValueOf(a) // 得到的是一个副本，要想得到本体，则传地址
	t := reflect.TypeOf(a)
	fmt.Println(v, t, v.Type(), v.Float(), v.Interface().(float64), int(v.Float()), v.Kind(), v.CanSet()) // kind种类（与Type类型有区别）

	var b MyTestReflectType
	b = 10
	v1 := reflect.ValueOf(&b)
	t1 := reflect.TypeOf(b)
	fmt.Println(v1, t1, v.Type(), v1.Interface(), v1.Kind(), v1.CanSet(), v1.Elem())
	v2 := v1.Elem() // 大概是获取反射的本体，此时可以用设置值了
	fmt.Println(v2.CanSet())
	v2.SetInt(12)
	fmt.Println(v2, v1.Elem())

	c := &MyTestReflectStruct{10, "嘻嘻"}
	rValue := reflect.ValueOf(c).Elem()
	rTyoe := rValue.Type()
	for i := 0; i < rValue.NumField(); i++ { // NumField和Field用于遍历结构体的反射
		f := rValue.Field(i) // 获取索引对象
		fmt.Printf("%d: %s %s = %v\n", i,
			rTyoe.Field(i).Name, f.Type(), f)
	}

	user := &MyTestReflectUser{
		Name: "user1", // 变量名小写时，表示未导出，反射不能访问
		Id:   10,
		Sex:  1,
	}
	va := reflect.ValueOf(user)
	vb := reflect.ValueOf(*user)
	fmt.Println(va.CanSet(), vb.CanSet(), va.Elem().CanSet()) // Elem()只能是得到指针的反射调用，否者会宕机
	fmt.Println(va.Elem().FieldByName("Name")) // FieldByName结构体反射独有，否则，其他使用，宕机
	newName := "user2"
	va.Elem().FieldByName("Name").Set(reflect.ValueOf(newName))
	fmt.Println(va.Elem().FieldByName("Name"))

	// 种类（Kind）与类型（Type）
	var ta MyTestReflectType
	ta = 10
	type newStruct struct {
		Type int `info:"11" value:"1"`
		// ``为标签，属于类型，不能赋值（键值只见才有空格，其他地方不能空格）
		// Get(key string)通过键获取值
		// Lookup(key string)通过键返回是否存在
	}
	fmt.Println(reflect.TypeOf(ta).Name(), reflect.TypeOf(ta).Kind())
	fmt.Println(reflect.TypeOf(newStruct{}).Name(), reflect.TypeOf(newStruct{}).Kind())
	nS := newStruct{Type: 1}
	if item, ok := reflect.TypeOf(nS).FieldByName("Type"); ok { //Tag属于类型
		fmt.Println(item.Tag.Get("info"))
	}
	// 控制判断 IsNil()是否为nil值 IsValid()是否有效
	fmt.Println("nil is valid: ", reflect.ValueOf(nil).IsValid())
	fmt.Println("*int IsValid: ", reflect.ValueOf((*int)(nil)).Elem().IsValid())
	fmt.Println("*int is nil", reflect.ValueOf((*int)(nil)).IsNil())

	// 已知反射类型是，可动态创建该类型的实例，获取的是指针类型
	taType := reflect.TypeOf(ta)
	newTa := reflect.New(taType)
	fmt.Println(newTa.Elem().Interface())
	newTa.Elem().Set(reflect.ValueOf(MyTestReflectType(12)))
	fmt.Println(newTa.Elem().Interface())

	// 反射调用函数(传入参数传入[]reflect.Value类型，返回[]reflect.Value类型)
	reflectFun := reflect.ValueOf(reflectTestCallFunc)
	argeList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(12)}
	result := reflectFun.Call(argeList)
	fmt.Println(result[0])
}

/**
 各种内置包
 */
func packageFunc()  {
	// os包（）进行系统的基本操作，如文件操作，目录操作，执行命令，信号与中断，进程，系统状态等等
	fmt.Println(os.Args) // 命令参数（返回切片）

	// go-qrcode 生成二维码(https://github.com/skip2/go-qrcode 下载库文件)
	// 命令行安装 go get -u github.com/skip2/go-qrcode/...

	// 获取二维码图片

	// func WriteFile(content string, level RecoveryLevel, size int, filename string) error
	// content 表示要生成二维码的内容，可以是任意字符串。
	// level 表示二维码的容错级别，取值有 Low、Medium、High、Highest。
	// size 表示生成图片的 width 和 height，像素单位。
	// filename 表示生成的文件名路径。
	// RecoveryLevel 类型其实是个 int，它的定义和常量如下。
	//qrcode.WriteFile("http://59.110.221.90:20001", qrcode.Medium, 256, "./qrcode.png")
	//
	//// 获取一个二维码bety的数据
	//betyData, err := qrcode.Encode("http://59.110.221.90:20001", qrcode.Medium, 256)
	//if err == nil {
	//	fmt.Println(betyData)
	//}
	//qr, err := qrcode.New("http://59.110.221.90:20001", qrcode.Medium)
	//if err == nil {
	//	qr.BackgroundColor = color.RGBA{25, 25, 25, 255}
	//	qr.ForegroundColor = color.RGBA{100, 100, 100, 255}
	//	qr.WriteFile(256, "./autoQrcode.png")
	//}

	// context包（上下文）
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()
	go contextFunc(ctx, 500 * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main: ", ctx.Err())
	}
}

func contextFunc(ctx context.Context, dur time.Duration)  {
	select {
	case <-ctx.Done():
		fmt.Println("handler: ", ctx.Err())
	case <-time.After(dur):
		fmt.Println("process request with", dur)
	}
}

/**
goroutine channel 并发
 */
func goroutineFunc()  {
	//c1 := make(chan int)  // 创建channel(只能用make创建)
	//c2 := make(chan string)
	c3 := make(chan interface{})

	times := 0
	go func() { // 此函数属于另一个goroutine
		for  {
			times++
			fmt.Print(times)
			time.Sleep(time.Second)
		}
	}()
	go func() {
		time.Sleep(time.Second)
		c3 <- "hell" // 使用<-符号发发送数据（发送将持续阻塞直到数据被接收）
	}()
	getData := <- c3 // 这这个goroutine中接受来自c3通道的数据（未接收到数据前此goroutine会阻塞）
	//<- c3 // 另一种写法（此种写法不需要用到通道的数据）
	fmt.Print(getData)
	//var input string
	//fmt.Scanln(&input)
	//runtime.Gosched()  // 暂停当前goroutine
	//cdan = make(<- chan int) // 单向通道（只能接受）
	//cdan2 = make(chan <- int) // 单向通道（只能发送）
	close(c3) // 关闭通道（关闭的通道不能发送数据，但依然可以内部访问）
	_, ok := <- c3 // ok返回false表示已经成功关闭
	//fmt.Print(ok)
	if ok == false {
		fmt.Print("通道关闭成功\n")
	}

	hcChan := make(chan int, 3) // 带缓冲通道（限制了发送的长度）
	// 此过程不阻塞goroutine
	hcChan <- 1
	hcChan <- 2
	hcChan <- 3
	fmt.Print("结尾\n")

	// 超时机制（通过select实现）
	//chan1 := make(chan int)
	//chan2 := make(chan int)
	//select {
	//case <-chan1:
	//// 如果chan1成功读到数据，则进行该case处理语句
	//case chan2 <- 1:
	//// 如果成功向chan2写入数据，则进行该case处理语句
	//default:
	//	// 如果上面都没有成功，则进入default处理流程
	//}
	chOutTime := make(chan bool)
	chNum := make(chan int)
	go func() {
		for  {
			select {
			case data := <- chNum:
				fmt.Printf("num = %d\n", data)
			case <- time.After(3 * time.Second): // time.After(一段时间之后，返回一个chan)
				fmt.Print("超时了")
				chOutTime <- true
			}
		}
	}()
	for i := 1; i< 10; i++ {
		chNum <- i
		time.Sleep(time.Second)
	}
	<- chOutTime
	fmt.Print("结束")
	// 不同goroutine对同一资源进行读写时，会发生竞争关系，需要锁住资源

	// 获取cpu核心数
	cpuNum := runtime.NumCPU()
	fmt.Print("cpu核心数：", cpuNum)
	runtime.GOMAXPROCS(16) // 设置开启的cpu核心数
}

func goroutineTest()  {
	ch := make(chan int)
	goroutine1 := func(c chan int) {
		for  {
			data := <- c
			fmt.Print("接收到的值", data)
			if data == 0 {
				break
			}
		}
		fmt.Print("goroutine1 结束")
		c <- 404
	}
	go goroutine1(ch)
	for i := 1; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	ch <- 0
	<- ch // 接受goroutine1发送的通道值
	fmt.Print("main结束")
}

func goroutineFunc2()  {
	cn := make(chan int)
	//cn2 := make(chan int)
	player := func(name string, ch chan int) {
		isDiu := rand.Intn(100) // 采用随机数判断是否丢球
		ball := <- ch
		if res := isDiu % 13; res == 0 {
			fmt.Printf("%s输了", name)
			//close(ch)
			ch <- 0
		} else {
			fmt.Printf("%s接球： %d", name, ball)
			ball ++
			ch <- ball
		}
	}
	go player("玩家1", cn)
	go player("玩家2", cn)
	cn <- 1
	<- cn
}

// 用通道模拟客户端服务端通信
func Cline(ch chan interface{}, rep string) (interface{}, error) {
	ch <- rep
	select {
	case data := <- ch:
		return data, nil
	case <- time.After(3 * time.Second):
		return "", errors.New("超时")
	}
}

func Server_Chan(ch chan interface{})  {
	for  {
		data := <- ch
		fmt.Print("服务端收到的数据：", data)
		time.Sleep(time.Second * 4)
		ch <- "回复"
	}
}
func goroutineFunc3()  {
	ch := make(chan interface{})
	go Server_Chan(ch)

	data, err := Cline(ch, "请求数据")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("收到回复了：", data)
	}
}

// chan模拟计时器

// 延迟回调
func goroutineFunc4()  {
	ch := make(chan int)
	fmt.Print("等待别人做完。。\n")
	time.AfterFunc(time.Second * 2, func() { // 等待一段时间做，第二个参数会回调
		fmt.Print("我做完了\n")
		ch <- 0
	})
	<- ch
	fmt.Print("该我了")
}

// 定点计时
func goroutineFunc5()  {
	tik := time.NewTicker(time.Millisecond * 500) // 打点器（每隔多少时间触发一次）
	timer := time.NewTimer(time.Second * 2) // 定时器（多少时间触发）
	count := 0
	for  {
		select {
		case <- timer.C: // （C成员返回的通知通道）
			goto stop
		case <- tik.C:
			count++
		}
	}
	stop:
		fmt.Print("over, count: ", count)
}

// 避免资源竞争的方式
func goroutineFunc6()  {
	var seq int64
	getId := func() int64 {
		// 这种写法会发生资源竞争（go run -race test.go运行会宕机）
		//atomic.AddInt64(&seq, 1)
		//return seq

		// 这种写法可以避免资源竞争
		return atomic.AddInt64(&seq, 1) // 原子操作函数（会将seq递增加第二个参数）
	}
	for i := 0; i < 10; i++ {
		go getId()
	}
	fmt.Print(getId())
}
// 互斥锁（sync.Mutex）与读写互斥锁（sync.RWMutex）
func goroutineFunc7()  {
	// 在上锁的情况下，其他goroutine会阻塞知道解锁
	var(
		a int
		aLock sync.Mutex // 互斥锁
	)
	setA := func(n int) {
		// 如果其他goroutine的aLock处于Lock状态，则此goroutine的Lock阻塞，直至其他goroutine的aLock处于Unlock状态
		aLock.Lock()
		defer aLock.Unlock()
		a = n
	}
	getA := func() int {
		aLock.Lock()
		defer aLock.Unlock()
		return a
	}
	setA(10)
	fmt.Print(getA())

	var (
		b int
		bLock sync.RWMutex // 读写互斥锁（适用于读多少写）
	)
	getB := func() int {
		bLock.RLock()
		defer bLock.RUnlock()
		return b
	}
	b = 10
	fmt.Print(getB())

}
// 等待组（sync.WaitGroup）
func goroutineFunc8()  {
	var waitGroup sync.WaitGroup

	urlList := []string{
		"www.baidu.com",
		"www.cocos.com",
		"www.go.com",
	}
	for _, value := range urlList {
		waitGroup.Add(1) // 等待组加1
		go func(str string) {
			time.Sleep(time.Second)
			fmt.Printf("%s\n", str)
			waitGroup.Done() // 等待组减1
		}(value)
	}
	waitGroup.Wait() // 等待计数不为0，则阻塞，等待计数为0，则取消阻塞
	fmt.Print("所有的goroutine都完了")
}

// 死锁（争夺资源而卡死）
func goroutineFunc9()  {
	type value struct {
		val int
		valLock sync.Mutex
	}
	var waitGroup sync.WaitGroup
	sum := func(v1, v2 *value) int {
		v1.valLock.Lock()
		defer v1.valLock.Unlock()
		time.Sleep(time.Second)
		v2.valLock.Lock() // 休眠1秒锁定v2，但在product中的v2已经上锁，等待v2解锁
		defer v2.valLock.Unlock()
		fmt.Print(v1.val + v2.val)
		waitGroup.Done()
		return v1.val + v2.val
	}
	product := func(v1, v2 *value) int {
		//waitGroup.Add(1)
		v2.valLock.Lock()
		defer v2.valLock.Unlock()
		time.Sleep(time.Second)
		v1.valLock.Lock() // 休眠1秒锁定v1，但在sum中的v1已经上锁，等待v1解锁
		defer v1.valLock.Unlock()
		fmt.Print(v1.val * v2.val)
		waitGroup.Done()
		return v1.val * v2.val
	}
	var val1, val2 value
	val1.val = 1
	val2.val = 2
	waitGroup.Add(1)
	waitGroup.Add(1)
	go sum(&val1, &val2)
	go product(&val1, &val2)
	fmt.Print("\n等待")
	waitGroup.Wait()
}

// 原子函数加锁
func yuanZiFunc()  {
	
}

// 文件处理

// json文件的读取
func jsonFileFunc()  {
	type Website struct {
		Name string
		Url string
	}
	ch := make(chan int)
	go func(cn chan int) {
		info := []Website{
			{
				Name: "Golang",
				Url: "www.golang.com",
			},
			{
				Name: "cocos",
				Url: "www.cocos.com",
			},
		}
		file, err := os.Create("test.json")
		if err != nil {
			fmt.Println("创建文件失败：", err.Error())
			cn <- 1
			return
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		err = encoder.Encode(info)
		if err != nil {
			fmt.Println("文件加密错误：", err.Error())
			cn <- 1
		} else {
			fmt.Print("文件加密成功")
			cn <- 2
		}
	}(ch)
	res := <- ch
	if res == 1 {

	} else {
		file, err := os.Open("./test.json")
		if err != nil {
			fmt.Println("打开文件失败：", err.Error())
			return
		}
		var infos []Website
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&infos)
		if err != nil {
			fmt.Println("解密失败：", err.Error())
		} else {
			fmt.Println("解密成功：", infos)
		}
	}

}

// xml文件的读写
func xmlFileFunc()  {
	//xml 包还支持更为复杂的标签，包括嵌套。例如标签名为 'xml:"Books>Author"'
	// 产生的是 <Books><Author>content</Author></Books> 这样的 XML 内容。同时除了
	// 'xml:", attr"' 之外，该包还支持 'xml:",chardata"' 这样的标签表示将该字段当做字符数据来写，
	// 支持 'xml:",innerxml"' 这样的标签表示按照字面量来写该字段，以及 'xml:",comment"'
	// 这样的标签表示将该字段当做 XML 注释。因此，
	// 通过使用标签化的结构体，我们可以充分利用好这些方便的编码解码函数，
	// 同时合理控制如何读写 XML 数据。
	type Website struct { // 命名首字母要大写
		Name string `xml:"name,attr"` // xml格式特有的标签
		Url string
		ChStr string
	}
	info := Website{
		Name:  "golang",
		Url:   "www.golang.com",
		ChStr: "golang学习",
	}
	ch := make(chan int)
	go func(cn chan int) {
		file, err := os.Create("./test.xml")
		if err != nil {
			fmt.Println("创建文件出错", err.Error())
			ch <- 1
			return
		}

		defer file.Close()
		fmt.Println(info)
		encode := xml.NewEncoder(file)
		err = encode.Encode(info)
		if err != nil {
			fmt.Println("加密错误：", err.Error())
			ch <- 1
		} else {
			fmt.Print("加密成功! \n")
			ch <- 2
		}
	}(ch)
	res := <- ch
	if res == 2 {
		file, err := os.Open("./test.xml")
		if err != nil {
			fmt.Println("打开文件出错：", err.Error())
			return
		}
		defer file.Close()

		var xmlInfo Website
		decode := xml.NewDecoder(file)
		err = decode.Decode(&xmlInfo)
		if err != nil {
			fmt.Println("解密失败：", err.Error())
		} else {
			fmt.Println("解密成功：", xmlInfo)
		}
	}
}

// Gob文件的读写
// Gob文件：go语言独有的二进制形式序列与反序列的数据格式（用于纯go环境）
func gobFileFunc()  {
	//info := map[string]string{
	//	"name": "golang",
	//	"url": "www.golang.com",
	//}
	ch := make(chan int)
	fileName := "./test.gob"
	fileName = "./test.txt"
	fileName = "./test.zip"
	go func(cn chan int) {
		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 077)
		if err != nil {
			fmt.Println("创建问题出错：", err.Error())
			ch <- 1
			return
		}
		defer file.Close()
		// gob文件的写入
		//encode := gob.NewEncoder(file)
		//if err = encode.Encode(info); err != nil {
		//	fmt.Println("写入文件失败：", err.Error())
		//	ch <- 1
		//} else {
		//	fmt.Println("写入文件成功")
		//	ch <- 2
		//}
		// txt文件的写入
		//write := bufio.NewWriter(file) // 创建一个带缓冲的写入器
		//str := "www.golang.com\t"
		//write.WriteString(str) // 写入缓冲
		//write.Flush() // 缓冲里写入文件里
		//ch <- 2

		// zip归档文件的写入
		buf := new(bytes.Buffer) // 创建一个缓冲区保护压缩内容
		w := zip.NewWriter(buf) // 创建一个压缩文档
		//defer w.Close()
		info := []struct{
			Name, Url string
		} {
			{"golang.txt", "www.golang.com"},
			{"cocos.txt", "www.cocos.com"},
		}
		for _, val := range info{
			fmt.Println("写入一个文件：", val)
			files, err := w.Create(val.Name) // 压缩文档去创建文件
			if err != nil {
				fmt.Println("创建文件出错：", err.Error())
			}
			_, err = files.Write([]byte(val.Url))
			if err != nil {
				fmt.Println("写入出错：", err.Error())
			}
		}
		err = w.Close() // 关闭压缩文件
		if err != nil {
			fmt.Println("关闭压缩文件出错：", err.Error())
		}
		// 讲压缩文档写入文件
		buf.WriteTo(file)
		ch <- 2
	}(ch)

	res := <- ch
	if res == 2 {
		//file, err := os.Open(fileName)
		file, err := zip.OpenReader(fileName) // zip文件打开
		defer file.Close()
		if err != nil {
			fmt.Println("文件打开失败: ", err.Error())
		} else {

			// gob文件的读取
			//var gobInfo map[string]string
			//decode := gob.NewDecoder(file)
			//if err = decode.Decode(&gobInfo); err != nil {
			//	fmt.Println("文件读取失败：", err.Error())
			//} else {
			//	fmt.Println("文件读取成功：", gobInfo)
			//}
			// txt文件的读取
			//reader := bufio.NewReader(file) // 创建一个读取器
			//for  {
			//	str, err := reader.ReadString('\t') // 读到一个空格就结束
			//	if err == io.EOF {  // 读完了
			//		break
			//	}
			//	fmt.Println("读到的数据：", str)
			//}
			//fmt.Println("读完了")
			// zip归档文件的读取
			for _, f := range file.File{ // 迭代压缩文件里的文件
				fmt.Printf("文件名%s\n", f.Name)
				rc, err := f.Open()
				if err != nil {
					fmt.Println("打开文件出错：", err.Error())
				}
				_, err = io.CopyN(os.Stdout, rc, int64(f.UncompressedSize64)) // 打印文件内容
				if err != nil {
					fmt.Println("读取文件内容出错：", err.Error())
				}
				rc.Close()
			}
		}
	}

}

/***************************实例******************************/

/**
利用结构体解析json数据
*/
type Screens struct {
	Size       float32 // 屏幕尺寸
	ResX, ResY int     // 屏幕水平和垂直分辨率
}
type Battery struct {
	Capacity int
}

func getJsonData() []byte {
	raw := &struct {
		Screens
		Battery
		HasTouchID bool
	}{
		Screens: Screens{ // 注意：此时字段名如果是小写，则初始化不成功
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			20,
		},
		HasTouchID: true,
	}
	jsonData, _ := json.Marshal(raw)
	fmt.Println(string(jsonData))
	return jsonData
}
func jsonTestFunc()  {
	jsonData := getJsonData()
	screnc := struct {
		Screens
		HasTouchID bool
	}{}
	json.Unmarshal(jsonData, &screnc) // 筛选json数据
	fmt.Println(screnc)
	battery := struct {
		Battery
		HasTouchID bool
	}{}
	json.Unmarshal(jsonData, &battery)
	fmt.Println(battery)
}

/**
 使用结构体实现单链表
 */
type Student struct {
	Name string
	Id int
}
type Node struct {
	Student
	Next *Node // 存放下一个Node的地址
}

func (head *Node) Creat () *Node {
	head = nil
	return head
}
func (p *Node)PrintLink()  {
	for p != nil {
		fmt.Println(p.Name, p.Id)
		p = p.Next
	}
}
func (newNode *Node) Insert (head *Node) *Node {
	if head == nil {
		fmt.Println("11")
		head = newNode
		head.Next = nil
	} else {
		headTem := head
		var lastHead *Node
		for headTem.Next != nil && headTem.Id < newNode.Id {
			lastHead = headTem
			headTem = headTem.Next
		}
		if headTem.Id >= newNode.Id {
			if lastHead == nil { // 此时元素只有一个
				head = newNode
				head.Next = headTem
				headTem.Next = nil
			} else {
				lastHead.Next = newNode
				newNode.Next = headTem
			}
		} else {
			headTem.Next = newNode
			newNode.Next = nil
		}
	}
	return head
}
func (head *Node) Delete (delNode *Node) *Node {
	headTem := head
	var lastHead *Node
	for headTem.Next != nil {
		if headTem.Id == delNode.Id && headTem.Name == delNode.Name {
			if lastHead == nil {
				head = head.Next
			} else {
				lastHead.Next = headTem.Next
			}
			goto onEnd
		}
		lastHead = headTem
		headTem = headTem.Next
	}
	if headTem.Next == nil {
		if headTem.Id == delNode.Id && headTem.Name == delNode.Name {
			if lastHead != nil {
				lastHead.Next = nil
			}
		}
	}
	onEnd:
		return head
}
func testListFunc()  {
	var newStudentList *Node
	sut1 := Node{Student{Name: "lihua", Id: 1001}, nil}
	sut2 := Node{Student{Name: "lili", Id: 1002}, nil}
	sut3 := Node{Student{Name: "haha", Id: 1003}, nil}
	sut4 := Node{Student{Name: "haha2", Id: 1006}, nil}
	sut5 := Node{Student{Name: "haha2", Id: 1004}, nil}
	newStudentList = newStudentList.Creat()
	newStudentList = sut3.Insert(newStudentList)
	newStudentList = sut1.Insert(newStudentList)
	newStudentList = sut4.Insert(newStudentList)
	newStudentList = sut2.Insert(newStudentList)
	newStudentList = sut5.Insert(newStudentList)
	newStudentList.PrintLink()
	fmt.Println("删除后")
	newStudentList = newStudentList.Delete(&sut2)
	newStudentList.PrintLink()
}

/*
讲一个字符串倒着输出
 */
func changeString(str string) string {
	var sArr [] string
	for _, value := range str{
		sArr = append([] string{string(value)}, sArr[:]...)
	}
	newStr := ""
	for _, value := range sArr{
		newStr = newStr + value
	}
	fmt.Println(newStr)
	return newStr
}

/**
 日志写入器
 */
func testLogFunc()  {
	logWriter := customLog.NewLogger()
	fWriter := customLog.NewFileWriter()
	consoleWrite := customLog.NewConsoleWriter()
	err := fWriter.SetFile("log")
	if err == nil {
		logWriter.RegisterWriter(fWriter)
	}
	logWriter.RegisterWriter(consoleWrite)
	logWriter.Log("测试一哈哈")
}

/**
 排序（内置sort方法，需要实现sort接口的三个方法： Len() int，Less(i, j int) bool， Swap(i, j int)）
 */
type myString []string

func (m myString) Len () int { // 获取长度的方法
	return len(m)
}

func (m myString) Less (i, j int) bool { // 比较的方法
	return m[i] < m[j]
}

func (m myString)Swap(i, j int)  { // 交换的方法
	m[i], m[j] = m[j], m[i]
}

// 排序结构体
type StudentInfo struct {
	Id int
	Name string
}

type StudentList []*StudentInfo

func (s StudentList) Len () int {
	return len(s)
}
func (s StudentList) Less (i, j int) bool {
	return s[i].Id < s[j].Id
}
func (s StudentList)Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func testSortFunc()  {
	strList := myString{
		"123",
		"234",
		"abc",
		"1",
		"2",
	}
	sort.Sort(strList) // 内置接口已实现，可以用内置sort方法了
	for _, v := range strList{
		fmt.Println(v)
	}
	// 可用内置的排序函数
	strList2 := []string{
		"1",
		"2",
		"5",
		"3",
	}
	sort.Strings(strList2) // 字符串排序（sort.Ints(a []int) 整型排序, sort.Float64s(a []Float64) 浮点数排序）
	for _, v := range strList2{
		fmt.Println(v)
	}

	stduentList := StudentList{
		&StudentInfo{
			Id:   123,
			Name: "user1",
		},
		&StudentInfo{
			Id:   128,
			Name: "user2",
		},
		&StudentInfo{
			Id:   124,
			Name: "user3",
		},
		&StudentInfo{
			Id:   125,
			Name: "user4",
		},
	}
	sort.Sort(stduentList) // 结构体
	for _, v := range stduentList{
		fmt.Printf("%v", v)
	}
	sort.Slice(stduentList, func(i, j int) bool { // 1.8以上才支持，直接对切片排序，传入比较函数
		return stduentList[i].Id > stduentList[j].Id
	})
	for _, v := range stduentList{
		fmt.Printf("%v", v)
	}
}

/**
 利用空接口实现字典
 */
func DictionaryFunc()  {
	dic := MyDictionary.NewDictionary()
	dic.Set("one", 123)
	dic.Set(2, "two")
	dic.Set(2, "three")
	dic.Set(3, "three")
	dic.Set("3", "four")
	fmt.Println("获取值：", dic.Get("3"))
	dic.Visit(func(i, j interface{}) bool {
		fmt.Println(i, j)
		return true
	})
}

/**
 利用类型断言type-switch
 */
type Walker interface {
	Walk()
}
type Flyer interface {
	Fly()
}
type Human struct {
	
}
func (h *Human) Walk ()  {
	fmt.Println("行走")
}
type Bird struct {

}

func (b *Bird) Fly ()  {
	fmt.Println("飞行")
}

func checkTypeFunc(data interface{})  {
	switch data.(type) { // data必须为interface类型
	case Walker:
		fmt.Println("这是行走类")
	case Flyer:
		fmt.Println("这是飞行类的")
	}
}

func typeSwitchTestFunc()  {
	var human interface{}
	human = new(Human)
	checkTypeFunc(human)
	bird := new(Bird)
	checkTypeFunc(bird)

	// 错误的类型
	_, err := os.Open("test")
	fmt.Println(err.(*os.PathError).Err)
	fmt.Println(os.IsExist(err)) // 文件已存储
	fmt.Println(os.IsNotExist(err)) // 文件未找到
	fmt.Println(os.IsPermission(err)) // 没有权限

	//var a interface{}
	////var b *Aaa
	//a = &Aaa{}
	//fmt.Println("11",a.(*Aaa))
}

//type Aaa struct {
//	a int
//}

/**
 迷你的web服务器
 */
func webServerFunc()  {
	TestWebServer.TWSHandlerFunc()
}
/*
音乐播放器
 */
func musicPlayerFunc()  {
	musicManager := SimpleMediaPlayer.NewMusicManager()
	fmt.Println(musicManager)
	music1 := &SimpleMediaPlayer.Music{
		Name:   "music1",
		Artist: "artist1",
		Source: "/music",
		Type:   "mp3",
	}
	music2 := &SimpleMediaPlayer.Music{
		Name:   "music2",
		Artist: "artist2",
		Source: "/music",
		Type:   "mp3",
	}
	musicManager.Add(music1)
	musicManager.Add(music2)
	fmt.Println(musicManager)
	fmt.Println(musicManager.Get("music1"))
	fmt.Println(musicManager.Len())
	fmt.Println(musicManager.Remove("music1"))
	fmt.Println(musicManager)

	SimpleMediaPlayer.MusicPlay()

}

/**
 有限状态机
 */

func FiniteStateMachineFunc()  {
	manager := FiniteStateMachine.GetStateManager()
	manager.Change = func(from, to FiniteStateMachine.State) {
		fmt.Printf("%s ---> %s\n\n", FiniteStateMachine.StateName(from), FiniteStateMachine.StateName(to))
	}
	var err error
	err = manager.Add(new(FiniteStateMachine.Idle))
	if err != nil {
		fmt.Print(err)
	}
	err = manager.Add(new(FiniteStateMachine.Move))
	if err != nil {
		fmt.Print(err)
	}
	err = manager.Add(new(FiniteStateMachine.Jump))
	if err != nil {
		fmt.Print(err)
	}

	err = manager.Transit("Idle")
	if err != nil {
		fmt.Print(err)
	}
	err = manager.Transit("Idle")
	if err != nil {
		fmt.Print(err)
	}
	err = manager.Transit("Move")
	if err != nil {
		fmt.Print(err)
	}
	err = manager.Transit("Move")
	if err != nil {
		fmt.Print(err)
	}
	err = manager.Transit("Jump")
	if err != nil {
		fmt.Print(err)
	}
	err = manager.Transit("Jump")
	if err != nil {
		fmt.Print(err)
	}
	err = manager.Transit("Idle")
	if err != nil {
		fmt.Print(err)
	}
	err = manager.Transit("Jump")
	if err != nil {
		fmt.Print(err)
	}
}
/**
 二叉树
 */
func BinaryTreeTest()  {
	root := BinaryTree.NewNode()
	root.SetData("root node")

	leftNode := BinaryTree.NewNode()
	leftNode.SetData("left node")

	rightNode := BinaryTree.NewNode()
	rightNode.SetData("right node")

	node1 := BinaryTree.NewNode()
	node1.SetData("node1")

	node2 := BinaryTree.NewNode()
	node2.SetData("node2")

	leftNode.Left = node1
	leftNode.Right = node2

	root.Left = leftNode
	root.Right = rightNode

	root.PrintBT()
	fmt.Printf("root depth: %v root LeafCount: %v", root.Depth(), root.LeafCount())
}

/**
 设计模式
 */
func DesignPatternFun()  {
	class1 := Factory.Create("class1")
	class1.Do()
}

/**
 制作gif动画
 */
func createGifAni ()  {
	palette := []color.Color{color.White, color.Black}
	const whiteIndex = 0
	const blackIndex = 1
	rand.Seed(time.Now().UTC().UnixNano())
	//if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, h *http.Request) {
			lissajous(w, palette, blackIndex)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	//}
}

func lissajous(out io.Writer, palette []color.Color, blackindex uint8)  {
	const (
		cycles = 5      //完整的x振荡器变化的个数
		res = 0.001     //角度分辨率
		size = 100      //图像画布包含[-size. .+size]
		nframes = 64    //动画中的帧数
		delay = 8       //以10ms为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 //y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase differenee
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackindex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //注意：忽略编码错误
}

/**
 Time包
 */
func timePackageFunc()  {
	t := time.Now()
	fmt.Print(t.Format("02 Jan 2006 15:04"))
	fmt.Printf("\n%04d,%02d,%04d", t.Year(), t.Month(), t.Day())

	fmt.Println(t)  // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 21.12.2011
	t = time.Now().UTC()
	fmt.Println(t)          // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	var week time.Duration = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)    // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC))  // Wed Dec 21 08:56:34 2011
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
}

/**
 客户管理系统
 */
func customerSystem()  {
	cView := View.GetView()
	cView.CustomerSystem()
}