// 获取 xml 所有节点名
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("")

	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	var t xml.Token
	var inItemGroup bool

	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			// 判断是否在 ItemGroup 中
			if inItemGroup {
				if name == "Compile" {
					fmt.Println(name)
				}
			} else {
				if name == "ItemGroup" {
					inItemGroup = true
				}
			}
			fmt.Println(name)
		case xml.EndElement:
			if inItemGroup {
				if token.Name.Local == "ItemGroup" {
					inItemGroup = false
				}
			}
		}
	}
}
