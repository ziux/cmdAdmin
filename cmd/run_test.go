package cmd

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {

	//os.Chdir("I:\\nginx-1.12.2")
	ch:= &Program{Name: "./nginx.exe",Arg: []string{"-t"},Dir: "I:\\nginx-1.12.2"}
	err := ch.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range ch.OutChan {
		fmt.Println(i)
	}


}
func TestOutStdout(t *testing.T) {

}
