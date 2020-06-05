package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput() (string,error){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println( "error:", err)
		return "", err
	}
	return scanner.Text(),nil
}

func SplitInput(data string)(string,[]string){
	list := strings.Split(data," ")
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