package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	var mykind string
	myparams := make([]interface{}, 20)
	kubectl := "kubectl"
	myexec := kubectl + "logs -f "

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

	if myexec != "" {
		myparams = []interface{}{"logs", "-f", userChoice}
		// fmt.Print("kubectl ")
		// fmt.Println(myparams)
		// out, err := ShellNew().Command("kubectl", myparams...).CombinedOutput()
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println(string(out))
		// }
	} else {
		fmt.Println("no command to run")
	}

	// copy to clipboard
	myparamstr := " "
	for _, v := range myparams {
		if str, ok := v.(string); ok {
			myparamstr += str + " "
		}
	}
	clipboard.WriteAll(userChoice)
	// text, _ := clipboard.ReadAll()
	fmt.Println(kubectl + myparamstr)
}
