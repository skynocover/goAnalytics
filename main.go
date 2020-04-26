package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	formate() //執行 go fmt

	docs, _ := filepath.Glob("*") //抓出所有的文件檔名
	for _, doc := range docs {
		if strings.HasSuffix(doc, ".go") { //如果是golang檔便建立物件
			var file tfile
			file.name = doc
			file.getObj() //抓出文件內所有的物件
			file.show()   //顯示物件內部的結構
			fmt.Println("")
		}
	}
}

func formate() { //執行go fmt .
	cmd := exec.Command("go", "fmt", ".") //檢視當前目錄下檔案
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
