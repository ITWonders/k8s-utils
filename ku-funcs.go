package main

import (
	"fmt"
	"regexp"

	sh "github.com/codeskyblue/go-sh"
)

func Kuget(kind string) ([]byte, error) {
	fmt.Println("Retriving kind: " + kind)
	session := sh.NewSession()
	mysession := session.Command("kubectl", "get", kind, "-o", "wide")

	// Note: below not working as once get the Output(), the stdout will be cleared.
	// subsequent Command(), and Output() will be empty
	// myout, err := mysession.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return myout, err
	// }
	// fmt.Println(string(myout))

	return mysession.Command("awk", `{ print $1 }`).
		Command("tail", "-n", "+2").
		Output()
}

func GetUserChoice(output string) string {
	// split the result based on newline
	re := regexp.MustCompile(`\n`)
	result := re.Split(output, -1)
	result = result[0 : len(result)-1]

	for k, v := range result {
		fmt.Printf("%d: %s\n", k, v)
	}

	fmt.Printf("\n\nChoose one number to ssh: ")
	var chosen int
	fmt.Scanf("%d", &chosen)

	return result[chosen]
}
