package main

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/urfave/cli/v2"
	"os"
)

func selfUpdate(slug string) error {
	selfupdate.EnableLog()

	previous := semver.MustParse(VERSION)
	latest, err := selfupdate.UpdateSelf(previous, slug)
	if err != nil {
		return err
	}

	if previous.Equals(latest.Version) {
		fmt.Println("Current binary is the latest version", VERSION)
	} else {
		fmt.Println("Update successfully done to version", latest.Version)
		fmt.Println("Release note:\n", latest.ReleaseNotes)
	}
	return nil
}

func CmdUpdate(c *cli.Context) error {
	slug := "nucleome/nucleserver"
	if err := selfUpdate(slug); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}
