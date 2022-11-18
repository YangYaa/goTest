package basic

import "fmt"

//定义了一个Printer接口，实现了Print方法的都可以称为Printer
type Printer interface {
	Print()
}

//佳能打印机
type CanonPrinter struct {
	printerContent string
}

//尼康打印机
type NikonPrinter struct {
	printerContent string
}

//佳能打印内容输出
func (printer CanonPrinter) Print() {
	fmt.Println(printer.printerContent, "print.")
}

//尼康打印内容输出
func (printer NikonPrinter) Print() {
	fmt.Println(printer.printerContent, "print.")
}

type PrintWoker struct {
	Printer
	name string
	age  int
}

func StructEmbedInterface() {
	canon := CanonPrinter{"This is canon"}
	nikon := NikonPrinter{"This is nikon"}
	canonWorker := PrintWoker{Printer: canon, name: "Zhang", age: 21}
	nikonWorker := PrintWoker{Printer: nikon, name: "Yang", age: 23}
	canonWorker.Printer.Print()
	nikonWorker.Printer.Print()
}
