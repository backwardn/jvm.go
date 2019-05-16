package main

import (
	"fmt"
)

// java [ options ] class [ arguments ]
type Command struct {
	Options Options
	Class   string
	Args    []string
}

func parseCommand(osArgs []string) (cmd *Command, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	args := osArgs[1:]
	cmd = &Command{
		Options: parseOptions(&args),
		Class:   removeFirst(&args),
		Args:    args,
	}

	return
}
