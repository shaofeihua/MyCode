/*
	打包编译多个静态文件
*/

package main

import (
	"embed"
	"io"
	_ "io/fs"
	"log"
	"os"
)

//go:embed version
//go.embed *.conf
var fs embed.FS

func main() {
	file, err := fs.Open("version")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(file)
	io.Copy(os.Stdout, file)

}
