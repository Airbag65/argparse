package example

import (
	"log"
	"os"

	"github.com/Airbag65/argparse"
)

//  =============================================================
//	# 	General interface to make CreateCmd work as intended	#
//	=============================================================
type Cmd interface {
	Execute()
}

//  =========================================================================
//	# 	Set up structs, and attach Execute function to match the interface	#
//	=========================================================================
type LsCmd struct{}

func (c *LsCmd) Execute() {
	// do stuff
}

type GetCmd struct {
	HasFlag   bool
	FlagValue string
}

func (c *GetCmd) Execute() {
	// do stuff
}

type NewCmd struct{}

func (c *NewCmd) Execute() {
	// do stuff
}

//  =====================================================
//	# 	Initialize the parser with commands and flags	#
//	=====================================================
func InitParser() *argparse.Parser {
	cmds := []string{
		"ls",
		"get",
		"new",
	}

	getFlagShort := argparse.NewFlag("-i", "Id of thing to get", true)
	getFlagLong := argparse.NewFlag("--id", "Id of thing to get", true)

	parser := argparse.New()

	for _, cmd := range cmds {
		if cmd == "get" {
			parser.AddCommand(cmd, argparse.AddFlag(getFlagShort), argparse.AddFlag(getFlagLong))
		} else {
			parser.AddCommand(cmd)
		}
	}
	return parser
}


//  =====================================================
//	# 	Use the parsed command to make create a command	#
//	=====================================================
func CreateCmd(pc *argparse.ParsedCommand) Cmd {
	switch pc.Command {
	case "ls":
		return &LsCmd{}
	case "get":
		return &GetCmd{
			HasFlag:   pc.Option != "",
			FlagValue: pc.Parameter,
		}
	case "new":
		return &NewCmd{}
	}

	return nil
}

//  =====================================================
//	# 	Representing the main function of the program	#
//	=====================================================
func Action() {
	p := InitParser()
	pc, err := p.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	c := CreateCmd(pc)
	c.Execute()
}
