package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	_1k = 1024
	_1m = _1k * _1k
	_1g = _1k * _1m
)

// java [ options ] class [ arguments ]
type Command struct {
	Options Options
	Class   string
	Args    []string
}

type Options struct {
	Classpath    string
	VerboseClass bool
	Xss          int
	Xcpuprofile  string
	XuseJavaHome bool
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

func parseOptions(args *[]string) Options {
	options := Options{
		Xss: 16 * _1k,
	}

	for hasMoreOptions(*args) {
		optionName := removeFirst(args)
		switch optionName {
		case "-cp", "-classpath":
			options.Classpath = removeFirst(args)
		case "-verbose", "-verbose:class":
			options.VerboseClass = true
		case "-Xcpuprofile":
			options.Xcpuprofile = removeFirst(args)
		case "-XuseJavaHome":
			options.XuseJavaHome = true
		default:
			if strings.HasPrefix(optionName, "-Xss") {
				options.Xss = parseXss(optionName)
			} else {
				panic("Unrecognized option: " + optionName)
			}
		}
	}

	return options
}

func hasMoreOptions(args []string) bool {
	return len(args) > 0 && args[0][0] == '-'
}

func removeFirst(args *[]string) string {
	first := (*args)[0]
	*args = (*args)[1:]
	return first
}

// -Xss<size>[g|G|m|M|k|K]
func parseXss(optionName string) int {
	size := optionName[4:]
	switch size[len(size)-1] {
	case 'g', 'G':
		return _1g * parseInt(size[:len(size)-1])
	case 'm', 'M':
		return _1m * parseInt(size[:len(size)-1])
	case 'k', 'K':
		return _1k * parseInt(size[:len(size)-1])
	default:
		return parseInt(size)
	}
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err.Error())
	}
	return i
}
