package main

import "fmt"

func main() {
	//定义map 并赋值
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	//定义空map 1
	m2 := make(map[string]int) // m2 == empty map
	//定义空map 2
	var m3 map[string]int // m3 == nil

	fmt.Println("m, m2, m3:")
	fmt.Println(m, m2, m3)
	//map[course:golang site:imooc quality:notbad name:ccmouse] map[] map[]

	fmt.Println("遍历 Map，Map的Key在输出时是无序的：")
	for k, v := range m {
		fmt.Println(k, v)
		//quality notbad
		//name ccmouse
		//course golang
		//site imooc
	}

	fmt.Println("获取指定key的value：")
	courseName := m["course"]
	fmt.Println(`m["course"] =`, courseName) //m["course"] = golang

	//试图读取map 中不存在的 key
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key 'cause' does not exist")
	}

	fmt.Println("删除指定key的键值对")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n",
		"name", name, ok)
	//m["name"] before delete: "ccmouse", true

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
	//m["name"] after delete: "", false
}