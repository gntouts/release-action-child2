package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"codeberg.org/msantos/embedexe/exec"
)

//go:embed bin
var bin []byte

var version string

func reverse(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}
func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "--version" {
		versionString := fmt.Sprintf("reversecho %s built with", version)
		fmt.Println(versionString)
		args = []string{"--version"}
	} else {
		finalArg := ""
		for _, arg := range args {
			finalArg = fmt.Sprintf("%s %s", finalArg, arg)
		}
		finalArg = reverse(finalArg)
		args = []string{finalArg}
	}

	cmd := exec.Command(bin, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalln("run:", cmd, err)
	}
}
