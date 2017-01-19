package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	// EnvDebug is environmental var to handle debug mode
	EnvDebug      = "GHR_DEBUG"
	EnvStackTrace = "GHR_TRACE"
)

// Exit codes are in value that represnet an exit code for a paticular error.
const (
	ExitCodeOK int = 0

	// Errors start at 10
	ExitCodeError = 10 + iota
	ExitCodeParseFlagsError
	ExitCodeBadArgs
	// ExitCodeInvalidURL
	// ExitCodeTokenNotFound
	// ExitCodeOwnerNotFound
	// ExitCodeRepoNotFound
	// ExitCodeRleaseError
)

const (
	defaultCheckTimeout = 2 * time.Second
	// defaultBaseURL      = "https://api.github.com/"
	// DefaultParallel     = -1
)

// Debugf prints debug output when EnvDebug is given
func Debugf(format string, args ...interface{}) {
	if env := os.Getenv(EnvDebug); len(env) != 0 {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {

	var (
		user    string
		decode  bool
		version bool
	)
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Usage = func() {
		fmt.Fprint(cli.errStream, helpText)
	}

	flags.StringVar(&user, "user", "", "")
	flags.StringVar(&user, "u", "", "")

	flags.BoolVar(&decode, "decode", false, "")
	flags.BoolVar(&decode, "d", false, "")

	flags.BoolVar(&version, "version", false, "")
	flags.BoolVar(&version, "v", false, "")

	// Parse flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if decode && user != "" {
		fmt.Fprintf(cli.outStream, OutputDecode(user))
		return ExitCodeOK
	}

	// Show version and check latest version release
	if version {
		fmt.Fprintf(cli.outStream, OutputVersion())
		return ExitCodeOK
	}

	if user == "" {
		fmt.Fprintf(cli.errStream, helpText)
		return ExitCodeOK
	}
	fmt.Fprintf(cli.outStream, OutputEncode(user))

	return ExitCodeOK
}

var helpText = `Usage: offer [-u]
Options:
  --decode, -d       Decode ID
  --version, -v      show version this app
`
