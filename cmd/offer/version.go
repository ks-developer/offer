package main

import (
	"bytes"
	"fmt"
)

// Name is application name
const Name = "offer"

// Version is application version
const Version string = "v0.0.1"

func OutputVersion() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s version %s", Name, Version)
	fmt.Fprintf(&buf, "\n")

	// Check latest version is release or not.
	// verCheckCh := make(chan *latest.CheckResponse)
	// go func() {
	// 	fixFunc := latest.DeleteFrontV()
	// 	githubTag := &latest.GithubTag{
	// 		Owner:             "ks-developer",
	// 		Repository:        "offer",
	// 		FixVersionStrFunc: fixFunc,
	// 	}

	// 	res, err := latest.Check(githubTag, fixFunc(Version))
	// 	if err != nil {
	// 		// Don't return error
	// 		Debugf("[ERROR] Check lastet version is failed: %s", err)
	// 		return
	// 	}
	// 	verCheckCh <- res
	// }()

	// select {
	// case <-time.After(defaultCheckTimeout):
	// case res := <-verCheckCh:
	// 	if res.Outdated {
	// 		fmt.Fprintf(&buf,
	// 			"Latest version of ghr is v%s, please upgrade!\n",
	// 			res.Current)
	// 	}
	// }

	return buf.String()
}
