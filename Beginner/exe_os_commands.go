package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("cmd", "/C", "git", "status").CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println("Command Successfully Executed")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>")
	output := string(out)
	fmt.Println(output)
	fmt.Println("<<<<<<<<<<<<<<<<<<<<")
}
