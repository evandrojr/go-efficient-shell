package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	execute "github.com/alexellis/go-execute/v2"
)

type Outcome struct {
	ExitCode int
	Stdout   string
	Stderr   string
}

func main() {
	// outcome := shellExecute("lssss -la", true)
	// outcome.ExitCode = 2

	cmdLongOutputFatal("ping 192.168.18.11")

}

func cmdLongOutputFatal(command string) {

	cmd := exec.Command("bash", "-c", command)

	// pipe the commands output to the applications
	// standard output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run still runs the command and waits for completion
	// but the output is instantly piped to Stdout
	if err := cmd.Run(); err != nil {
		log.Fatalf("Command: [%s] ERROR: [%s]\n", command, err)
	}
}

func printOutcome(outcome Outcome) {
	fmt.Printf("stdout:\n%s\nstderr:\n%s\nexit-code: %d\n", outcome.Stdout, outcome.Stderr, outcome.ExitCode)
}

func shellExecute(command string, showOutcome bool) Outcome {
	cmd := execute.ExecTask{
		Command:     "sh",
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
		printOutcome(outcome)
	}
	return outcome
}
