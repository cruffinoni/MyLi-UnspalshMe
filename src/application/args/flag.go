package args

import (
	"errors"
	"log"
	"os"
	"strconv"
)

const projectName = "unsplash-me"

var HelpMessage = errors.New("help message")

type ProgArg struct {
	Page  int64
	Query string
}

func printHelpMessage() {
	log.Printf("At least 1 argument is required and 2 is acceptable:\n"+
		"\t%v [QUERY] [optional: PAGE]\n"+
		"QUERY: The query against unsplash\n"+
		"PAGE: The page number", projectName)
}

func New() (ProgArg, error) {
	argsLen := len(os.Args)
	if argsLen < 2 || argsLen > 3 {
		printHelpMessage()
		return ProgArg{}, errors.New("invalid arg count")
	}
	args := ProgArg{
		Page:  1,
		Query: os.Args[1],
	}
	if args.Query == "-h" || args.Query == "--help" {
		printHelpMessage()
		return ProgArg{}, HelpMessage
	}
	if argsLen == 3 {
		var err error
		args.Page, err = strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			log.Print("can't parse page number\n")
			return ProgArg{}, err
		}
	}
	return args, nil
}
