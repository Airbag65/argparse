package argparse_test

import (
	"testing"

	"github.com/Airbag65/argparse"
)

func TestPass(t *testing.T) {
	str := "hello"
	if str != "hello" {
		t.Error("strings did not match")
	}
}

func TestCreateFlag(t *testing.T) {
	f := argparse.NewFlag("--test", "Test Flag", true)

	if f.Flag != "--test" {
		t.Error("flag name was not set")
	}

	if f.Description != "Test Flag" {
		t.Error("flag description was not set")
	}

	if !f.Forced {
		t.Error("flag.Forced was not set")
	}
}

func TestParse(t *testing.T) {
	p := argparse.New()
	p.AddCommand("ls")
	f := argparse.NewFlag("-t", "test flag", true)
	p.AddCommand("test", argparse.AddFlag(f))

	command, err := p.Parse([]string{"passport", "ls"})
	if err != nil {
		t.Error(err)
	}
	if command.Command != "ls" {
		t.Error("Command was not created correctly")
	}
	if command.Option != "" {
		t.Error("Command was not created correctly")
	}
	if command.Parameter != "" {
		t.Error("Command was not created correctly")
	}

	command, err = p.Parse([]string{"passport", "test", "-t", "data"})
	if err != nil {
		t.Error(err)
	}
	if command.Command != "test" {
		t.Error("Command was not created correctly")
	}
	if command.Option != "-t" {
		t.Error("Command was not created correctly")
	}
	if command.Parameter != "data" {
		t.Error("Command was not created correctly")
	}

	command, err = p.Parse([]string{"passport", "test", "-e", "data"})
	if err == nil {
		t.Error("correctly parsed incorrect command")
	}
}
