package main

import (
	"fmt"
)

// java [ options ] class [ arguments ]
type Command struct {
	Options *Options
	Class   string
	Args    []string
}

func ParseCommand(osArgs []string) (cmd *Command, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	argReader := &ArgReader{osArgs[1:]}
	cmd = &Command{
		Options: parseOptions(argReader),
		Class:   argReader.removeFirst(),
		Args:    argReader.args,
	}

	return
}

func PrintUsage() {
	fmt.Println("usage: jvmgo [-options] class [args...]")
}
