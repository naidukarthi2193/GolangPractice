package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	var allsentences []string
	for {
		fmt.Print("--> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)
		if text != "" {
			allsentences = append(allsentences, text)
		} else {
			fmt.Println(allsentences)
			break
		}

	}
}
