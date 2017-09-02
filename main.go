package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/albrow/stringset"
)

const usage = `Confirm is a tiny utility for confirming actions.

Usage:

    confirm [flags] message

'message' is a message that will be printed out to the user.

After the message has been printed out, by default confirm will wait to receive
either 'yes' or 'y' (case-insensitive) before continuing. You can enter 'no' or
'n' (case-insensitive) to cancel the action and return a non-zero status code.

The flags are:

    help
        Print usage information.

    case-sensitive
        Make confirmation of the given input case-sensitive instead of the
        default, case-insensitive.

    confirm-with
        Comma-separated list of values to indicate confirmation. If one of these
        values is provided, confirm will exit with a status code of 0. (default:
        'y,yes').

    cancel-with
        Comma-separated list of values to indicate cancelation. If one of these
        values is provided, confirm will indicate the action was canceled and
        exit with a status code of 1. An empty string means no value will result
        in cancelation. (default: 'n,no').

`

// maxAttempts is the maximum number of times to receive user input
// before giving up. If a user enters an invalid input maxAttempts
// times, we will exit with a non-zero status code.
const maxAttempts = 3

func main() {
	var (
		help          bool
		caseSensitive bool
		confirmWith   string
		cancelWith    string
	)
	flag.BoolVar(&help, "help", false, "")
	flag.BoolVar(&caseSensitive, "case-sensitive", false, "")
	flag.StringVar(&confirmWith, "confirm-with", "y,yes", "")
	flag.StringVar(&cancelWith, "cancel-with", "n,no", "")
	flag.Parse()

	if help {
		print(usage)
		os.Exit(0)
	}

	// Validate arguments.
	if len(flag.Args()) != 1 {
		print("Confirm requires exactly one argument: the message to be printed out.\n")
		os.Exit(1)
	}

	// Parse confirm-with and cancel-with flags.
	confirmSet := stringset.New()
	if confirmWith == "" {
		// It's okay for cancelSet to be empty, but if confirmSet is empty,
		// it's impossible to confirm. This is an error.
		print("Confirm requires at least one value for the confirm-with flag.\n")
		os.Exit(1)
	}
	confirmSet = stringset.NewFromSlice(strings.Split(confirmWith, ","))
	cancelSet := stringset.New()
	if cancelWith != "" {
		cancelSet = stringset.NewFromSlice(strings.Split(cancelWith, ","))
	}

	// Parse arguments.
	message := flag.Arg(0)
	if !strings.HasSuffix(message, "\n") {
		message += "\n"
	}
	print(message)

	// We run a loop to check for user input up to maxAttempts
	// times.
	attempts := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		// If caseSensitive is false, we want to treat the
		// input as case-insensitive.
		if !caseSensitive {
			input = strings.ToLower(input)
		}
		// If the set of accept values contains the input,
		// exit with a status code of 0.
		if confirmSet.Contains(input) {
			os.Exit(0)
		}
		// If the set of cancel values contains the input,
		// print a message indicating the action was canceled
		// and exit with a status code of 1.
		if cancelSet.Contains(input) {
			print("Action canceled by user.\n")
			os.Exit(1)
		}
		// If we exceeded the maximum number of attempts,
		// print a message and exit with a status code of 1.
		attempts++
		if attempts >= maxAttempts {
			print("Could not confirm action: Too many attempts with invalid input.\n")
			os.Exit(1)
		}
		print("Unexpected input.\n")
		print(message)
	}
	// Note: we don't expect to reach this line, but it's
	// here to make sure that we don't fail silently if
	// something does go wrong.
	print("Confirm could not read user input.\n")
	os.Exit(1)
}
