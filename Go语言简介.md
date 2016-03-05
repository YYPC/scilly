Go语言简介

作者：陈皓

转载有删节。

```
- Hello World
文件名 hello.go

package main //声明本文件的package名

import "fmt" //import语言的fmt库——用于输出

func main() {
    fmt.Println("hello world")
}

- 运行
你可以有两种运行方式，

解释执行（实际是编译成a.out再执行）

$go run hello.go
hello world

编译执行

$go build hello.go

$ls
hello hello.go

$./hello

hello world

- 自己的package

你可以使用GOPATH环境变量，或是使用相对路径来import你自己的package。

Go的规约是这样的：

1）在import中，你可以使用相对路径，如 ./或 ../ 来引用你的package

2）如果没有使用相对路径，那么，go会去找$GOPATH/src/目录。

使用相对路径: import "./haoel"  //import当前目录里haoel子目录里的所有的go文件

使用GOPATH路径: import "haoel"  //import 环境变量 $GOPATH/src/haoel子目录里的所有的go文件

fmt输出格式: fmt包和libc里的那堆使用printf， scanf，fprintf，fscanf 很相似。下面的东西对于C程序员不会陌生。

注意：Println不支持，Printf才支持%式的输出：

package main

import "fmt"
import "math"

func main() {
    fmt.Println("hello world")

    fmt.Printf("%t\n", 1==2)
    fmt.Printf("二进制：%b\n", 255)
    fmt.Printf("八进制：%o\n", 255)
    fmt.Printf("十六进制：%X\n", 255)
    fmt.Printf("十进制：%d\n", 255)
    fmt.Printf("浮点数：%f\n", math.Pi)
    fmt.Printf("字符串：%s\n", "hello world")
}

当然，也可以使用如\n\t\r这样的和C语言一样的控制字符

- 变量和常量

变量的声明很像 javascript，使用 var关键字。注意：go是静态类型的语言，下面是代码：

//声明初始化一个变量
var  x int = 100
var str string = "hello world"</pre>
//声明初始化多个变量
var  i, j, k int = 1, 2, 3

//不用指明类型，通过初始化值来推导
var b = true //bool型
还有一种定义变量的方式（这让我想到了Pascal语言，但完全不一样）

x := 100 //等价于 var x int = 100;

常量很简单，使用const关键字：

const s string = "hello world"
const pi float32 = 3.1415926

- 数组
直接看代码（注意其中的for语句，和C很相似吧，就是没有括号了）

func main() {
    var a [5]int
    fmt.Println("array a:", a)

    a[1] = 10
    a[3] = 30
    fmt.Println("assign:", a)

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("init:", b)

    var c [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            c[i][j] = i + j
        }
    }
    fmt.Println("2d: ", c)
}

运行结果：

array a: [0 0 0 0 0]
assign: [0 10 0 30 0]
len: 5
init: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]

-数组的切片操作
这个很Python了。

a := [5]int{1, 2, 3, 4, 5}

b := a[2:4] // a[2] 和 a[3]，但不包括a[4]
fmt.Println(b)

b = a[:4] // 从 a[0]到a[4]，但不包括a[4]
fmt.Println(b)

b = a[2:] // 从 a[2]到a[4]，且包括a[2]
fmt.Println(b)

- 分支循环语句
1. if语句

注意：if 语句没有圆括号，而必需要有花括号

//if 语句
if x % 2 == 0 {
    //...
}
//if - else
if x % 2 == 0 {
    //偶数...
} else {
    //奇数...
}

//多分支
if num < 0 {
    //负数
} else if num == 0 {
    //零
} else {
    //正数
}

1. switch 语句

注意：switch语句没有break，还可以使用逗号case多个值

switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    case 4,5,6:
        fmt.Println("four, five, six")
    default:
        fmt.Println("invalid value!")
}

1. for 语句

前面你已见过了，下面再来看看for的三种形式：（注意：Go语言中没有while）

//经典的for语句 init; condition; post
for i := 0; i<10; i++{
     fmt.Println(i)
}

//精简的for语句 condition
i := 1
for i<10 {
    fmt.Println(i)
    i++
}

//死循环的for语句 相当于for(;;)
i :=1
for {
    if i>10 {
        break
    }
    i++
}

- 关于分号
从上面的代码我们可以看到代码里没有分号。其实，和C一样，Go的正式的语法使用分号来终止语句。和C不同的是，这些分号由词法分析器在扫描源代码过程中使用简单的规则自动插入分号，因此输入源代码多数时候就不需要分号了。

规则是这样的：如果在一个新行前方的最后一个标记是一个标识符（包括像int和float64这样的单词）、一个基本的如数值这样的文字、或以下标记中的一个时，会自动插入分号：

break continue fallthrough return ++ -- ) }
通常Go程序仅在for循环语句中使用分号，以此来分开初始化器、条件和增量单元。如果你在一行中写多个语句，也需要用分号分开。

注意：无论任何时候，你都不应该将一个控制结构（(if、for、switch或select）的左大括号放在下一行。如果这样做，将会在大括号的前方插入一个分号，这可能导致出现不想要的结果。

- map
map在别的语言里可能叫哈希表或叫dict，下面是和map的相关操作的代码，代码很容易懂

func main(){
    m := make(map[string]int) //使用make创建一个空的map

    m["one"] = 1
    m["two"] = 2
    m["three"] = 3

    fmt.Println(m) //输出 map[three:3 two:2 one:1] (顺序在运行时可能不一样)
    fmt.Println(len(m)) //输出 3

    v := m["two"] //从map里取值
    fmt.Println(v) // 输出 2

    delete(m, "two")
    fmt.Println(m) //输出 map[three:3 one:1]

    m1 := map[string]int{"one": 1, "two": 2, "three": 3}
    fmt.Println(m1) //输出 map[two:2 three:3 one:1] (顺序在运行时可能不一样)

    for key, val := range m1{
        fmt.Printf("%s => %d \n", key, val)
        /*输出：(顺序在运行时可能不一样)
            three => 3
            one => 1
            two => 2*/
    }
}

- 指针
Go语言一样有指针，看代码

var i int = 1
var pInt *int = &i
//输出：i=1     pInt=0xf8400371b0       *pInt=1
fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

*pInt = 2
//输出：i=2     pInt=0xf8400371b0       *pInt=2
fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

i = 3
//输出：i=3     pInt=0xf8400371b0       *pInt=3

fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)


- 内存分配
Go具有两个分配内存的机制，分别是内建的函数new和make。他们所做的事不同，所应用到的类型也不同，这可能引起混淆，但规则却很简单。

new 是一个分配内存的内建函数，但不同于其他语言中同名的new所作的工作，它只是将内存清零，而不是初始化内存。new(T)为一个类型为T的新项目分配了值为零的存储空间并返回其地址，也就是一个类型为*T的值。用Go的术语来说，就是它返回了一个指向新分配的类型为T的零值的指针。

make(T, args)函数的目的与new(T)不同。它仅用于创建切片、map和chan（消息管道），并返回类型T（不是*T）的一个被初始化了的（不是零）实例。这种差别的出现是由于这三种类型实质上是对在使用前必须进行初始化的数据结构的引用。例如，切片是一个具有三项内容的描述符，包括指向数据（在一个数组内部）的指针、长度以及容量，在这三项内容被初始化之前，切片值为nil。对于切片、映射和信道，make初始化了其内部的数据结构并准备了将要使用的值。如：

下面的代码分配了一个整型数组，长度为10，容量为100，并返回前10个数组的切片

make([]int, 10, 100)

以下示例说明了new和make的不同。

var p *[]int = new([]int)   // 为切片结构分配内存；*p == nil；很少使用
var v  []int = make([]int, 10) // 切片v现在是对一个新的有10个整数的数组的引用

// 不必要地使问题复杂化：
var p *[]int = new([]int)
fmt.Println(p) //输出：&[]
*p = make([]int, 10, 10)
fmt.Println(p) //输出：&[0 0 0 0 0 0 0 0 0 0]
fmt.Println((*p)[2]) //输出： 0

// 习惯用法:

v := make([]int, 10)
fmt.Println(v) //输出：[0 0 0 0 0 0 0 0 0 0]

- 函数
老实说，我对Go语言这种反过来声明变量类型和函数返回值的做法有点不满（保持和C一样的不可以吗? 呵呵）

package main
import "fmt"

func max(a int, b int) int { //注意参数和返回值是怎么声明的

    if a > b {
        return a
    }
    return b
}

func main(){
    fmt.Println(max(4, 5))
}

- 函数返回多个值

Go中很多Package 都会返回两个值，一个是正常值，一个是错误，如下所示：

package main
import "fmt"

func main(){
    v, e := multi_ret("one")
    fmt.Println(v,e) //输出 1 true

    v, e = multi_ret("four")
    fmt.Println(v,e) //输出 0 false

    //通常的用法(注意分号后有e)
    if v, e = multi_ret("four"); e {
        // 正常返回
    }else{
        // 出错返回
    }
}

func multi_ret(key string) (int, bool){
    m := map[string]int{"one": 1, "two": 2, "three": 3}

    var err bool
    var val int

    val, err = m[key]

    return val, err
}

- 函数不定参数

例子很清楚了，我就不多说了

func sum(nums ...int) {
    fmt.Print(nums, " ")  //输出如 [1, 2, 3] 之类的数组
    total := 0
    for _, num := range nums { //要的是值而不是下标
        total += num
    }
    fmt.Println(total)
}
func main() {
    sum(1, 2)
    sum(1, 2, 3)

    //传数组
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}

- 函数闭包

nextNum这个函数返回了一个匿名函数，这个匿名函数记住了nextNum中i+j的值，并改变了i,j的值，于是形成了一个闭包的用法

func nextNum() func() int {
    i,j := 1,1
    return func() int {
        var tmp = i+j
        i, j = j, tmp
        return tmp
    }
}
//main函数中是对nextNum的调用，其主要是打出下一个斐波拉契数
func main(){
    nextNumFunc := nextNum()
    for i:=0; i<10; i++ {
        fmt.Println(nextNumFunc())
    }
}

- 函数的递归

和c基本是一样的

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))
}

- 结构体
Go的结构体和C的基本上一样，不过在初始化时有些不一样，Go支持带名字的初始化。

type Person struct {
    name string
    age  int
    email string
}

func main() {
    //初始化
    person := Person{"Tom", 30, "tom@gmail.com"}
    person = Person{name:"Tom", age: 30, email:"tom@gmail.com"}

    fmt.Println(person) //输出 {Tom 30 tom@gmail.com}

    pPerson := &person

    fmt.Println(pPerson) //输出 &{Tom 30 tom@gmail.com}

    pPerson.age = 40
    person.name = "Jerry"
    fmt.Println(person) //输出 {Jerry 40 tom@gmail.com}
}

- 结构体方法
不多说了，看代码吧。

注意：Go语言中没有public, protected, private的关键字，所以，如果你想让一个方法可以被别的包访问的话，你需要把这个方法的第一个字母大写。这是一种约定。

type rect struct {
    width, height int
}

func (r *rect) area() int { //求面积
    return r.width * r.height
}

func (r *rect) perimeter() int{ //求周长
    return 2*(r.width + r.height)
}

func main() {
    r := rect{width: 10, height: 15}

    fmt.Println("面积: ", r.area())
    fmt.Println("周长: ", r.perimeter())

    rp := &r
    fmt.Println("面积: ", rp.area())
    fmt.Println("周长: ", rp.perimeter())
}

- 接口和多态
接口意味着多态，下面是一个经典的例子，不用多说了，自己看代码吧。

//---------- 接 口 --------//
type shape interface {
    area() float64 //计算面积
    perimeter() float64 //计算周长
}

//--------- 长方形 ----------//
type rect struct {
    width, height float64
}

func (r *rect) area() float64 { //面积
    return r.width * r.height
}

func (r *rect) perimeter() float64 { //周长
    return 2*(r.width + r.height)
}

//----------- 圆  形 ----------//
type circle struct {
    radius float64
}

func (c *circle) area() float64 { //面积
    return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 { //周长
    return 2 * math.Pi * c.radius
}

// ----------- 接口的使用 -----------//
func interface_test() {
    r := rect {width:2.9, height:4.8}
    c := circle {radius:4.3}

    s := []shape{&r, &c} //通过指针实现

    for _, sh := range s {
        fmt.Println(sh)
        fmt.Println(sh.area())
        fmt.Println(sh.perimeter())
    }
}

- 错误处理 – Error接口
函数错误返回可能是C/C++时最让人纠结的东西的，Go的多值返回可以让我们更容易的返回错误，其可以在返回一个常规的返回值之外，还能轻易地返回一个详细的错误描述。通常情况下，错误的类型是error，它有一个内建的接口。

type error interface {
    Error() string
}

还是看个示例吧：

package main

import "fmt"
import "errors"

//自定义的出错结构
type myError struct {
    arg  int
    errMsg string
}
//实现Error接口
func (e *myError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.errMsg)
}

//两种出错
func error_test(arg int) (int, error) {
    if arg < 0  {
         return -1, errors.New("Bad Arguments - negtive!")
     }else if arg >256 {
        return -1, &myError{arg, "Bad Arguments - too large!"}
    }
    return arg*arg, nil
}

//相关的测试
func main() {
    for _, i := range []int{-1, 4, 1000} {
        if r, e := error_test(i); e != nil {
            fmt.Println("failed:", e)
        } else {
            fmt.Println("success:", r)
        }
    }
}

程序运行后输出：

failed: Bad Arguments - negtive!
success: 16
failed: 1000 - Bad Arguments - too large!

- 错误处理 – Defer
下面的程序对于每一个熟悉C语言的人来说都不陌生（有资源泄露的问题），C++使用RAII来解决这种问题。

func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }

    written, err = io.Copy(dst, src)
    dst.Close()
    src.Close()
    return
}

Go语言引入了Defer来确保那些被打开的文件能被关闭。如下所示：（这种解决方式还是比较优雅的）

func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}

Go的defer语句预设一个函数调用（延期的函数），该调用在函数执行defer返回时立刻运行。该方法显得不同常规，但却是处理上述情况很有效，无论函数怎样返回，都必须进行资源释放。

我们再来看一个defer函数的示例：

for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}

被延期的函数以后进先出（LIFO）的顺行执行，因此以上代码在返回时将打印4 3 2 1 0。

总之，我个人觉得defer的函数行为有点怪异，我现在还没有完全搞清楚。

- 错误处理 – Panic/Recover
对于不可恢复的错误，Go提供了一个内建的panic函数，它将创建一个运行时错误并使程序停止（相当暴力）。该函数接收一个任意类型（往往是字符串）作为程序死亡时要打印的东西。当编译器在函数的结尾处检查到一个panic时，就会停止进行常规的return语句检查。

下面的仅仅是一个示例。实际的库函数应避免panic。如果问题可以容忍，最好是让事情继续下去而不是终止整个程序。

var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}

当panic被调用时，它将立即停止当前函数的执行并开始逐级解开函数堆栈，同时运行所有被defer的函数。如果这种解开达到堆栈的顶端，程序就死亡了。但是，也可以使用内建的recover函数来重新获得Go程的控制权并恢复正常的执行。 对recover的调用会通知解开堆栈并返回传递到panic的参量。由于仅在解开期间运行的代码处在被defer的函数之内，recover仅在被延期的函数内部才是有用的。

你可以简单地理解为recover就是用来捕捉Painc的，防止程序一下子就挂掉了。

下面是一个例程，很简单了，不解释了

func g(i int) {
    if i>1 {
        fmt.Println("Panic!")
        panic(fmt.Sprintf("%v", i))
    }

}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()

    for i := 0; i < 4; i++ {
        fmt.Println("Calling g with ", i)
        g(i)
        fmt.Println("Returned normally from g.")
     }
}

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

运行结果如下：（我们可以看到Painc后的for循环就没有往下执行了，但是main的程序还在往下走）

Calling g with  0
Returned normally from g.
Calling g with  1
Returned normally from g.
Calling g with  2
Panic!
Recovered in f 2
Returned normally from f.


你习惯这种编程方式吗？我觉得有点诡异。呵呵。

好了，上面是一Go语言相关的编程语法的介绍，我没有事无巨细，只是让你了解一下Go语言是长什么样的。
