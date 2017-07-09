package main

import (
	"log"
	"os"

	gopen "github.com/petermbenjamin/go-open"
	"github.com/tj/docopt"
)

// Version is the package version
var Version = "0.0.1"

// Usage is the package infomation
const Usage = `
  Usage:
    qiita-home <username>
    qiita-home -h | --help
    qiita-home --version
`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	username := args["<username>"].(string)

	if username != "" {
		gopen.Open("https://qiita.com/" + username)
		os.Exit(0)
	}
}
