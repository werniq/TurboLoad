package logger

import (
	"log"
	"os"
)

var (
	ErrorLogger = log.New(os.Stdout, "[!] [ERROR]: \t", log.Lshortfile|log.Ldate|log.Ltime)
	InfoLogger  = log.New(os.Stdout, "[*] [INFO]: \t", log.Ldate|log.Ltime)
)
