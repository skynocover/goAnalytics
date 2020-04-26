package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type tfile struct { //文件的物件
	name    string    //文件的檔名
	objects []tobject //文件內物件的陣列
}

func (file *tfile) getObj() {
	var obj tobject
	check := false //開始抓取物件屬性的通道預設關閉

	content := Readfile(file.name)
	for _, value := range content {
		if strings.HasPrefix(value, "type") { //如果此行開頭是type則抓取物件名並開啟通道
			obj.getName(value)
			check = true
		} else if strings.HasPrefix(value, "}") && check == true { //若通道開啟且此行開頭為}則推進物件並重置物件的屬性,並關閉通道
			file.objects = append(file.objects, obj)
			obj.properties = []string{}
			check = false
		} else if check == true { //若通道開啟且開頭不為}則推進物件的屬性
			obj.getProp(value)
		}
	}

	for i := 0; i < len(file.objects); i++ { //每個物件根據名稱去抓取自己的function
		file.objects[i].getFunc(content)
	}
}

func (file *tfile) show() {
	namespace := prt(file.name, 0)
	for j := 0; j < len(file.objects); j++ {
		objspace := prt(file.objects[j].name, namespace)
		prt(strconv.Itoa(len(file.objects[j].properties))+" properties", objspace)
		for k := 0; k < len(file.objects[j].properties); k++ {
			prt(file.objects[j].properties[k], objspace)
		}
		prt(strconv.Itoa(len(file.objects[j].function))+" functions", objspace)
		for m := 0; m < len(file.objects[j].function); m++ {
			prt(file.objects[j].function[m], objspace)
		}
	}
}

func prt(input string, space int) int { //根據輸入的前置space印出空格,之後列印文字,回傳前置空格與輸入文字的一半格數
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(input)
	return strings.Count(input, "")/2 + space
}

func Readfile(filename string) (lines []string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines = strings.Split(string(content), "\r\n")
	if len(lines) == 1 { //如果不是以windowns的方式分行則只會有一段
		lines = strings.Split(string(content), "\n")
	}
	return
}
