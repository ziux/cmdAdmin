package util

import "strings"

func ParseCommand(command string)(string,[]string){
	list := strings.Split(command," ")
	name := ""
	var args []string
	if len(list) >= 1{
		name = list[0]
	}
	if len(list)>1{
		args = list[1:]
	}
	return name,args
}
