package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	var mykind string
	var myexec string

	// provided myexec to run
	if len(os.Args) == 2 {
		mykind = os.Args[1]
	} else if len(os.Args) > 2 {
		mykind = os.Args[1]
		myexec = os.Args[2]
	} else {
		mykind = "po"
	}

	output, err := Kuget(mykind)
	if err != nil {
		fmt.Println(err)
		return
	}

	userChoice := GetUserChoice(string(output))

	clipboard.WriteAll(userChoice)
	text, _ := clipboard.ReadAll()
	fmt.Println(text)

	if myexec != "" {
		fmt.Println(myexec)
	}
}
