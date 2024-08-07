package ges

import (
	"context"
	"fmt"

	execute "github.com/alexellis/go-execute/v2"
)

type Outcome struct {
	ExitCode int
	Stdout   string
	Stderr   string
}

type Shell struct {
	Name    string
	Path    string
	Command string
}

var shell Shell

func initializeShell() {
	if shell.Name == "" {
		shell = Shell{
			Name:    "sh",
			Path:    "/usr/bin/sh",
			Command: "sh",
		}
	}
}

func PrintOutcome(outcome Outcome) {
	fmt.Printf("stdout:\n%s\nstderr:\n%s\nexit-code: %d\n", outcome.Stdout, outcome.Stderr, outcome.ExitCode)
}

func ShellExec(command string, showOutcome bool) Outcome {
	initializeShell()
	cmd := execute.ExecTask{
		Command:     shell.Command,
		Args:        []string{"-c", command},
		StreamStdio: false,
	}

	res, err := cmd.Execute(context.Background())
	if err != nil {
		panic(err)
	}
	outcome := Outcome{
		res.ExitCode,
		res.Stdout,
		res.Stderr,
	}
	if showOutcome {
		fmt.Println("Command:", command)
		PrintOutcome(outcome)
	}
	return outcome
}
