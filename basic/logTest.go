package basic

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Info = log.New(infoHandle, "INFO: ", log.LstdFlags|log.Lshortfile|log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func CateLog() {
	//Print different types of logs to different files
	file, err := os.OpenFile("/home/test.log", os.O_APPEND|os.O_WRONLY|os.O_APPEND, 666)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	defer file.Close()
	Init(file, os.Stdout, os.Stderr)
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
func Log() {
	//config log file path
	logFile, err := os.OpenFile("./testLogger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate) //config logger
	log.SetPrefix("[sctp:]")                                    //config log prefix
	log.Println("This is a test log")
	v := "predefine string"
	log.Printf("This is a %s log。\n", v)
	log.Fatalln("This will trigger fatal")
	log.Panicln("This will trigger panic")
}
