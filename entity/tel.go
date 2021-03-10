package entity

type Tel struct {
	Mobile     string // 电话
	Msg        string // Unicode消息
	SerialPort string // 串口名称
	Baud       string // 波特率
}

func NewTel() *Tel {
	return &Tel{
		Mobile:     "",
		Msg:        "",
		SerialPort: "/dev/ttyUSB0",
		Baud:       "115200",
	}
}
