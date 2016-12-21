package main

import (
	"fmt"
	"regexp"

	sh "github.com/codeskyblue/go-sh"
)

// ShellNew create new shell session and ready to take new command
func ShellNew() *sh.Session {
	return sh.NewSession()
}

// Kuget exec "kubectl get xxx"
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

// GetUserChoice returns choice of user
func GetUserChoice(output string) string {
	// split the result based on newline
	re := regexp.MustCompile(`\n`)
	result := re.Split(output, -1)
	result = result[0 : len(result)-1]

	for k, v := range result {
		fmt.Printf("%d: %s\n", k, v)
	}

	fmt.Printf("\n\nChoose one number (will be copied): ")
	var chosen int
	fmt.Scanf("%d", &chosen)

	return result[chosen]
}
