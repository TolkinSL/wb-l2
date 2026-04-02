package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		fmt.Print(": ")

		if !scanner.Scan() {
			fmt.Println("\nexit")
			return
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		go handleSignal(sigChan)

		if err := runLine(line); err != nil {
			fmt.Println("error:", err)
		}
	}
}

func handleSignal(sigChan chan os.Signal) {
	<-sigChan
	fmt.Println("\nInterrupted")
}

func runLine(line string) error {
	parts := strings.Split(line, "|")

	var cmds []*exec.Cmd

	for _, part := range parts {
		args := strings.Fields(strings.TrimSpace(part))
		if len(args) == 0 {
			continue
		}

		if isBuiltin(args[0]) {
			return runBuiltin(args)
		}

		cmds = append(cmds, exec.Command(args[0], args[1:]...))
	}

	for i := 0; i < len(cmds)-1; i++ {
		out, err := cmds[i].StdoutPipe()
		if err != nil {
			return err
		}
		cmds[i+1].Stdin = out
		cmds[i].Stderr = os.Stderr
	}

	last := cmds[len(cmds)-1]
	last.Stdout = os.Stdout
	last.Stderr = os.Stderr

	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return err
		}
	}

	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return err
		}
	}

	return nil
}

func isBuiltin(cmd string) bool {
	switch cmd {
	case "cd", "pwd", "echo", "kill", "ps":
		return true
	}
	return false
}

func runBuiltin(args []string) error {
	switch args[0] {

	case "cd":
		if len(args) < 2 {
			return fmt.Errorf("cd: missing path")
		}
		return os.Chdir(args[1])

	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(dir)

	case "echo":
		fmt.Println(strings.Join(args[1:], " "))

	case "kill":
		if len(args) < 2 {
			return fmt.Errorf("kill: missing pid")
		}
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			return err
		}
		return proc.Signal(syscall.SIGKILL)

	case "ps":
		cmd := exec.Command("ps", "aux")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	return nil
}
