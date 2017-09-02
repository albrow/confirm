package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

const usage = `Confirm is a tiny utility for confirming actions.

Usage:

    confirm [flags] message

'message' is a message that will be printed out to the user.

After the message has been printed out, confirm will wait to
receive either 'yes' or 'y' (case-insensitive) before
continuing. You can enter 'no' or 'n' (case-insensitive) to
cancel the action and return a non-zero status code.

The fags are:

    --help
        Print usage information.
    --case-sensitive
        Make confirmation of the given input case-sensitive instead of the default, case-insensitive.
`

// maxAttempts is the maximum number of times to receive user input
// before giving up. If a user enters an invalid input maxAttempts
// times, we will exit with a non-zero status code.
const maxAttempts = 3

func main() {
	var (
		help          bool
		caseSensitive bool
	)
	flag.BoolVar(&help, "help", false, "Print usage information.")
	flag.BoolVar(&caseSensitive, "case-sensitive", false, "Make confirmation of the given input case-sensitive instead of the default, case-insensitive.")
	flag.Parse()

	if help {
		print(usage)
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		print("Confirm requires exactly one argument, the message to be printed out.\n")
		os.Exit(1)
	}

	message := flag.Arg(0)
	if !strings.HasSuffix(message, "\n") {
		message += "\n"
	}
	print(message)

	attempts := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		// If caseSensitive is false, we want to treat the
		// input as case-insensitive.
		if !caseSensitive {
			input = strings.ToLower(input)
		}
		switch input {
		case "y", "yes":
			os.Exit(0)
		case "n", "no":
			print("Action cancelled by user.\n")
			os.Exit(1)
		default:
			attempts++
			if attempts >= maxAttempts {
				print("Could not confirm action: Too many attempts with invalid input.\n")
				os.Exit(1)
			}
			print("Unexpected input. Please try again.\n")
			print(message)
		}
	}
	print("Confirm could not read user input.\n")
	os.Exit(1)
}
