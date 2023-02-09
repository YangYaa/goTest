package panicRecover

/*
Recover 是一个Go语言的内建函数
可以让进入宕机流程中的 goroutine 恢复过来
recover 仅在延迟函数 defer 中有效
在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果
如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行。

通常来说，不应该对进入 panic 宕机的程序做任何处理
但有时，需要我们可以从宕机中恢复，至少我们可以在程序崩溃前，做一些操作。
如当 web 服务器遇到不可预料的严重问题时，在崩溃前应该将所有的连接关闭。
如果不做任何处理，会使得客户端一直处于等待状态。
如果 web 服务器还在开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助调试。

Go语言没有异常系统。其使用 panic 触发宕机类似于其他语言的抛出异常，recover 的宕机恢复机制就对应其他语言中的 try/catch 机制
*/
import (
	"fmt"
	"runtime"
)

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}

//该函数传入一个匿名函数或闭包后的执行函数。
//当传入函数以任何形式发生 panic 崩溃后,可以将崩溃发生的错误打印出来。
//同时允许后面的代码继续运行，不会造成整个进程的崩溃。

func ProtectRun(entry func()) {
	// 延迟处理的函数
	//--panic 和 recover 的组合有如下特性：
	//---有 panic 没 recover，程序宕机!。
	//---有 panic 也有 recover，程序不会宕机，执行完对应的 defer 后，从宕机点退出当前函数后继续执行。
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()

	entry()
}
func PanicTest() {
	fmt.Println("运行前")
	// 允许一段手动触发的错误
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		// 使用panic传递上下文
		panic(&panicContext{
			"手动触发panic",
		})
		fmt.Println("手动宕机后")
	})
	// 故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
}
