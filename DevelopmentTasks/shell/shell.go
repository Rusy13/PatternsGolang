package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" || text == "quit" {
			break
		}

		parts := strings.Fields(text)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "cd":
			if len(parts) > 1 {
				os.Chdir(parts[1])
			}
		case "pwd":
			dir, _ := os.Getwd()
			fmt.Println(dir)
		case "echo":
			fmt.Println(strings.Join(parts[1:], " "))
		case "kill":
			if len(parts) > 1 {
				exec.Command("kill", parts[1]).Run()
			}
		case "ps":
			exec.Command("ps").Run()
		default:
			cmd := exec.Command(parts[0], parts[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
	}
}
