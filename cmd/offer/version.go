package main

import (
	"bytes"
	"fmt"
)

// Name is application name
const Name = "offer"

// Version is application version
const Version string = "v0.0.7"

func OutputVersion() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s version %s", Name, Version)
	fmt.Fprintf(&buf, "\n")

	return buf.String()
}
