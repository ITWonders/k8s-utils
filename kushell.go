package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

var kind string
var mycommand []string

func main() {
	kind = "po"
	mycommand = []string{"kubectl", "exec", "-it"}

	if len(os.Args) > 1 {
		kind = os.Args[1]
		mycommand = []string{"gcloud", "compute", "ssh"}
		gcloudssh()
	} else {
		kushell()
	}
}

func kuget() ([]byte, error) {
	session := sh.NewSession()
	return session.Command("kubectl", "get", kind, "-o", "wide").
		Command("awk", `{ print $1 }`).
		Command("tail", "-n", "+2").
		Output()
}

func getUserChoice(output string) string {
	// split the result based on newline
	re := regexp.MustCompile(`\n`)
	result := re.Split(output, -1)

	for k, v := range result {
		fmt.Printf("%d: %s\n", k, v)
	}

	var chosen int
	fmt.Scanf("%d", &chosen)

	return result[chosen]
}

func kushell() {
	lines, err := kuget()

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	selected := getUserChoice(string(lines))
	mycommand := append(mycommand, selected, "--", "/bin/sh")

	fmt.Println(strings.Join(mycommand, " "))
	binary, _ := exec.LookPath("kubectl")
	syscall.Exec(
		binary,
		mycommand,
		os.Environ())
}

func gcloudssh() {
	lines, err := kuget()

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	selected := getUserChoice(string(lines))
	mycommand := append(mycommand, selected)

	fmt.Println(strings.Join(mycommand, " "))
	binary, _ := exec.LookPath("gcloud")
	syscall.Exec(
		binary,
		mycommand,
		os.Environ())

}
