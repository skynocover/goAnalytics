package main

import (
	"strings"
)

type tobject struct {
	name       string
	properties []string
	function   []string
}

func (obj *tobject) getName(input string) {
	str := strings.Split(input, "type")
	str = strings.Split(str[1], "struct")
	obj.name = strings.Trim(str[0], " ")
}

func (obj *tobject) getProp(input string) {
	str := strings.TrimSpace(input)
	arr := strings.Split(str, " ")
	obj.properties = append(obj.properties, arr[0])
}

func (obj *tobject) getFunc(input []string) {
	for i := 0; i < len(input); i++ {
		if strings.HasPrefix(input[i], "func") && strings.Contains(input[i], "*"+obj.name) {
			str := strings.Split(input[i], ")")
			str = strings.Split(str[1], "(")
			obj.function = append(obj.function, strings.TrimSpace(str[0]))
		}
	}
}
