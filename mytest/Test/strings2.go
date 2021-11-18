// 将 struct 序列化为 xml，然后进行反序列化
package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	// attr 是以属性的方式转换：<person name="xiaonan">
	// 双引号里面第一个字符串是相应字段的昵称，如下所示 Name 和 name
	// Name string `xml:"name,attr"`
	Name string
	Age  int
}

func main() {
	p := person{Name: "xiaonan", Age: 18}
	var data []byte
	var err error

	// 将 struct 序列化为 xml 。使用 xml.Marshal 转化出来的是连续的字符串，不包含格式，即： <person><Name>xiaonan</Name><Age>18</Age></person>
	// 如果想将序列化后的数据格式化输出来，则要使用 xml.MarshalIndent 。多出两个参数分别是前缀和缩进。以下代码前缀为空，缩进为空格
	// if data, err := xml.Marshal(p); err != nil {
	if data, err = xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	p2 := new(person)
	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n", p2)
}
