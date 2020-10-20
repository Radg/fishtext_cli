package main

import (
	"errors"
	"flag"
	"os"
)

const (
	messageDefault = "No commands provided. Usage examples: \n" +
		"fishtext_cli sentence -count=3\n" +
		"fishtext_cli paragraph -count=5\n" +
		"fishtext_cli title -count=7"
	messageSentenceUsage = "fishtext_cli sentence -count=3"
	messageParagraphUsage = "fishtext_cli paragraph -count=3"
	messageTitleUsage = "fishtext_cli title -count=3"
)

func ParseArguments(osArgs []string) (command string, count int, err error) {
	switch osArgs[1] {
	case "sentence":
		sentenceCommand := flag.NewFlagSet("sentence", flag.ExitOnError)
		sentenceFlag := sentenceCommand.Int("count", 3, messageSentenceUsage)
		err := sentenceCommand.Parse(os.Args[2:])
		if err != nil {
			return "", 0, err
		}
		return "sentence", *sentenceFlag,nil
		
	case "paragraph":
		paragraphCommand := flag.NewFlagSet("paragraph", flag.ExitOnError)
		paragraphFlag := paragraphCommand.Int("count", 3, messageParagraphUsage)
		err := paragraphCommand.Parse(os.Args[2:])
		if err != nil {
			return "", 0, err
		}
		return "paragraph", *paragraphFlag, nil

	case "title":
		titleCommand := flag.NewFlagSet("title", flag.ExitOnError)
		titleFlag := titleCommand.Int("count", 1, messageTitleUsage)
		err := titleCommand.Parse(os.Args[2:])
		if err != nil {
			return "", 0, err
		}
		return "title", *titleFlag, nil
	default:
		return "", 0, errors.New("unknown command given")
	}

}
