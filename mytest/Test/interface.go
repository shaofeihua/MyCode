package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connecter //嵌入式接口
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.name)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnected.")
}
func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}
