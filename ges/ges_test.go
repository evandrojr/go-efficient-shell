package ges

import "testing"

func TestHelloName(t *testing.T) {
	command := "env | grep SHELL"

	Exec(command, true)

}
