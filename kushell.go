package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"os"
	"os/exec"
	"regexp"
	"syscall"
)

func main() {
	kushell()
}

func kushell() {
	session := sh.NewSession()

	lines, err := session.Command("kubectl", "get", "po", "-o", "wide").
		Command("awk", `{ print $1 }`).
		Command("tail", "-n", "+2").
		Output()

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// split the result based on newline
	re := regexp.MustCompile(`\n`)
	result := re.Split(string(lines), -1)

	for k, v := range result {
		fmt.Printf("%d: %s\n", k, v)
	}

	var chosen int
	fmt.Scanf("%d", &chosen)

	binary, _ := exec.LookPath("kubectl")
	syscall.Exec(
		binary,
		[]string{"kubectl", "exec", "-it", result[chosen], "--", "/bin/sh"},
		os.Environ())
}
