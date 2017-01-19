package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun_encodeFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	command := "offer -u 1234"
	args := strings.Split(command, " ")

	if got, want := cli.Run(args), ExitCodeOK; got != want {
		t.Fatalf("%q exits %d, want %d", command, got, want)
	}

	want := fmt.Sprint("45474433DD1EA94D")
	got := outStream.String()
	if !strings.Contains(got, want) {
		t.Fatalf("%q output %q, want = %q", command, got, want)
	}
}

func TestRun_decodeFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	command := "offer -d -u 45474433DD1EA94D"
	args := strings.Split(command, " ")

	if got, want := cli.Run(args), ExitCodeOK; got != want {
		t.Fatalf("%q exits %d, want %d", command, got, want)
	}

	want := fmt.Sprint("1234")
	got := outStream.String()
	if !strings.Contains(got, want) {
		t.Fatalf("%q output %q, want = %q", command, got, want)
	}
}

// func TestRun_versionFlag(t *testing.T) {
// 	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
// 	cli := &CLI{outStream: outStream, errStream: errStream}
// 	command := "offer --version"
// 	args := strings.Split(command, " ")

// 	if got, want := cli.Run(args), ExitCodeOK; got != want {
// 		t.Fatalf("%q exits %d, want %d", command, got, want)
// 	}

// 	want := fmt.Sprintf("offer version %s", Version)
// 	got := outStream.String()
// 	if !strings.Contains(got, want) {
// 		t.Fatalf("%q output %q, want = %q", command, got, want)
// 	}
// }
