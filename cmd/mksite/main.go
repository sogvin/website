package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/sogvin/website"
)

func main() {
	var (
		cli          = cmdline.NewBasicParser()
		prefix       = cli.Option("-p, --prefix", "write pages to").String("./docs")
		showVersion  = cli.Flag("-v, --version")
		checkRelease = cli.Flag("-c, --check-release")
	)
	cli.Parse()

	log.SetFlags(0)

	switch {
	case showVersion:
		fmt.Println(website.Version())

	case checkRelease:
		if website.Version() == "unreleased" {
			log.Fatal("not ready for release, fix changelog")
		}

	default:
		os.MkdirAll(prefix, 0722)
		website := website.NewWebsite()
		if err := website.SaveTo(prefix); err != nil {
			log.Fatal(err)
		}
	}
}
