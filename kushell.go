package main

import (
	"fmt"
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

// check if a particular container/pods come with bash shell
func isBashNotExist(mystr string) bool {
	containers := []string{"ecv-go", "ecv-storage", "go-", "myfuse-"}

	for _, v := range containers {
		var re = regexp.MustCompile(v)
		if re.MatchString(mystr) {
			return true
		}
	}

	return false
}

func kushell() {
	lines, err := Kuget(kind)
	shelltype := "/bin/bash"

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	selected := GetUserChoice(string(lines))
	if isBashNotExist(selected) {
		shelltype = "/bin/sh"
	}
	mycommand = append(mycommand, selected, "--", shelltype)

	fmt.Println(strings.Join(mycommand, " ") + "\n")
	binary, _ := exec.LookPath("kubectl")
	syscall.Exec(
		binary,
		mycommand,
		os.Environ())
}

func gcloudssh() {
	lines, err := Kuget(kind)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	selected := GetUserChoice(string(lines))
	mycommand = append(mycommand, selected)

	fmt.Println(strings.Join(mycommand, " "))
	binary, _ := exec.LookPath("gcloud")
	syscall.Exec(
		binary,
		mycommand,
		os.Environ())

}
