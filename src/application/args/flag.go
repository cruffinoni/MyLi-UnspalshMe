package args

import (
	"errors"
	"log"
	"os"
	"strconv"
)

const projectName = "unsplash-me"

type ProgArg struct {
	Page  int64
	Query string
}

func New() (ProgArg, error) {
	argsLen := len(os.Args)
	if argsLen < 2 || argsLen > 3 {
		log.Printf("At least 1 argument is required and 2 is acceptable:\n"+
			"\t%v [QUERY] [PAGE]\n"+
			"QUERY: The query against unsplash"+
			"PAGE: The page number", projectName)
		return ProgArg{}, errors.New("invalid arg count")
	}
	args := ProgArg{
		Page:  1,
		Query: os.Args[1],
	}
	if argsLen == 3 {
		var err error
		args.Page, err = strconv.ParseInt(os.Args[2], 10, 8)
		if err != nil {
			log.Printf("can't parse page number (input '%v') to an int: %v\n", os.Args[2], err)
			return ProgArg{}, err
		}
	}
	return args, nil
}
