// sample project sample.go
package main

import (
	"fmt"
	"time"
	"unsafe"
)

func varDefine() {
	// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	var i int = 5
	// 第二种，根据值自行判定变量类型。
	var j = 5
	// 第三种，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：v_name := value
	intVal := 5

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(intVal)

}

func constVarDefine() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Println("面积为 : %d", area)
	println(a, b, c)

	// 常量还可以用作枚举：
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值
	const (
		m = "abc"
		n = len(m)
		p = unsafe.Sizeof(m)
	)

	println(m, n, p)
}

func iotaDefine() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

func swap(x, y string) (string, string) {
	return y, x
}

func funcMultiReturnVal() {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}

func arrayDefine() {
	var n [10]int /* n 是一个长度为 10 的数组 */

	var i, j int
	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

func ptrDefine() {
	var a int = 20    /* 声明实际变量 */
	var ip *int = nil /* 声明指针变量 */
	ip = &a           /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func structDefine() {
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.glodon.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.glodon.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.glodon.com"})

	var book Books
	book.author = "baijd"
	book.book_id = 123445
	book.subject = "new go run"
	book.title = "hell go"

	fmt.Printf("book title : %s\n", book.title)
	fmt.Printf("book author : %s\n", book.author)
	fmt.Printf("book subject : %s\n", book.subject)
	fmt.Printf("book book_id : %d\n", book.book_id)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
func sliceDefine() {
	// 1.切片定义
	var slice1 []int
	var slice2 []int = make([]int, 10) //slice2 := make([]type, len)
	// 也可以指定容量，其中capacity为可选参数。 make([]T, length, capacity)
	// 2.切片初始化
	// []表示是切片类型，{1,2,3, 4, 5}初始化值依次是1,2,3, 4, 5.其cap=len=5
	s1 := []int{1, 2, 3, 4, 5}
	// 初始化切片s,是数组arr的引用
	var arr [10]int
	var i int
	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		arr[i] = i + 100 /* 设置元素为 i + 100 */
	}
	s2 := arr[:]
	// 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	var startIndex, endIndex int
	startIndex = 1
	endIndex = 10
	s3 := arr[startIndex:endIndex]
	// 默认 endIndex 时将表示一直到arr的最后一个元素
	s4 := arr[startIndex:]
	// 默认 startIndex 时将表示从arr的第一个元素开始
	s5 := arr[startIndex:endIndex]
	// 通过切片s初始化切片s1
	// s6 := make([]int, len, cap) //通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
	s6 := make([]int, 3, 5)

	printSlice(slice1)
	printSlice(slice2)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)
	printSlice(s4)
	printSlice(s5)
	printSlice(s6)
	// 切片是可索引的，并且可以由 len() 方法获取长度。
	// 切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
	fmt.Printf("len=%d cap=%d slice=%v\n", len(s6), cap(s6), s6)

	// 一个切片在未初始化之前默认为 nil，长度为 0
	var s7 []int
	printSlice(s7)

	// 可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]
	s8 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", s8[1:4])

	// 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
	// 下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

// range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
// 在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
func rangeDefine() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func mapDefine() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")
	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func interfaceDefine() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func errorDefine() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func concurrentDefine() {
	go say("world")
	say("hello")
}

// 通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
// ch := make(chan int) 通道在使用前必须先创建

// ch <- v    // 把 v 发送到通道 ch
// v := <-ch  // 从 ch 接收数据 并把值赋给 v
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

// 注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。
func chanDefine() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

// 通道缓冲区 ch := make(chan int, 100)
// 带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
// 不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

// 注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
func chanBufferDefine() {
	ch := make(chan int, 2)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func foreachChan() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}

func basicOperate() {
	varDefine()
	constVarDefine()
	iotaDefine()
	funcMultiReturnVal()
	arrayDefine()
	ptrDefine()
	structDefine()
	sliceDefine()
	rangeDefine()
	mapDefine()
	interfaceDefine()
	errorDefine()
	concurrentDefine()
	chanDefine()
	chanBufferDefine()
	foreachChan()
}
