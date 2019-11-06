package main // 所属包（main包且含有main为命令源码文件）

import (
	"container/list"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"myPro/MyDictionary"
	"myPro/SimpleMediaPlayer"
	"myPro/TestWebServer"
	"myPro/customLog"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

func init()  {
	fmt.Println("初始化")
}

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
	musicPlayerFunc()
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