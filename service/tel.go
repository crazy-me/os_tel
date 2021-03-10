package service

import (
	"fmt"
	"github.com/crazy-me/os_tel/entity"
	"github.com/crazy-me/os_tel/utils"
	"github.com/tarm/serial"
	"log"
	"os"
	"strconv"
	"strings"
)

var l log.Logger

func init() {
	_ = utils.CreateDir("logs")
	logFile, _ := os.OpenFile("logs/tel_error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	l.SetOutput(logFile)
	l.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	l.SetPrefix("[os_tel]")
}

func Run(tel *entity.Tel) bool {
	l.Printf("%s It's not a valid phone number\n", tel.Mobile)
	if ok := utils.IsTel(tel.Mobile); !ok {
		l.Printf("tel:%s It's not a valid phone number\n", tel.Mobile)
		fmt.Printf("tel:%s It's not a valid phone number\n", tel.Mobile)
		return false
	}
	v, _ := strconv.Atoi(tel.Baud)
	c := &serial.Config{Name: tel.SerialPort, Baud: v}
	serialPort, err := serial.OpenPort(c)
	if err != nil {
		l.Println(err)
		return false
	}
	defer serialPort.Close()
	ascii := strconv.QuoteToASCII(tel.Msg)
	telMsg := fmt.Sprintf("%s, %s", tel.Mobile, strings.Replace(ascii[1:len(ascii)-1], "\\u", "", -1))
	fmt.Println(telMsg)
	writeLen, err := serialPort.Write([]byte(telMsg))
	if err != nil || 0 == writeLen {
		l.Printf("%s, Write fail", telMsg, err)
		return false
	}

	buf := make([]byte, 1024)
	readLen, err := serialPort.Read(buf)
	if err != nil {
		l.Printf("%s, Read fail", telMsg, err)
		return false
	}

	fmt.Printf("read:%q\n", buf[:readLen])
	return true
}
