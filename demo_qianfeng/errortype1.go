package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("www.baidu.com")
	//fmt.Println(err)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("addr:", addr) // [110.242.68.3 110.242.68.4]

	addr1, err1 := net.LookupHost("www.shdjaishfi.com")
	fmt.Println("err1:", err1) // lookup www.shdjaishfi.com: no such host
	//if err1 != nil {
	//	fmt.Println(err1) // lookup www.shdjaishfi.com: no such host
	//}
	if ins, ok := err1.(*net.DNSError); ok {
		if ins.Timeout() {
			fmt.Println("解析超时...")
		} else if ins.Temporary() {
			fmt.Println("临时性错误...")
		} else {
			fmt.Println("通用性错误...")
		}
	}
	fmt.Println("addr1:", addr1) // []
}
