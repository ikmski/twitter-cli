package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

var (
	version  string
	revision string
)

var config globalConfig
var auth authentication

const (
	configFileName         = "config.toml"
	authenticationFileName = "authentication.toml"
)

func getConfigFilePath() string {

	filePath := ""
	isExist := false
	curDir, err := os.Getwd()
	if err == nil {
		filePath = filepath.Join(curDir, configFileName)
		_, err = os.Stat(filePath)
		if err == nil {
			isExist = true
		}
	}

	if !isExist {
		filePath = filepath.Join(os.Getenv("HOME"), ".config", "twitter-cli", configFileName)
	}

	return filePath
}

func cmdTimeline(c *cli.Context) error {

	t := newTwitter()

	t.userStream()

	return nil
}

func cmdSearch(c *cli.Context) error {

	if c.NArg() < 1 {
		return fmt.Errorf("argument is required")
	}
	arg := c.Args().Get(0)

	t := newTwitter()

	t.publicStream(arg)

	return nil
}

func main() {

	configFile := getConfigFilePath()
	_, err := os.Stat(configFile)
	if err != nil {

		fmt.Println(err)
		os.Exit(1)

	} else {

		_, err := toml.DecodeFile(getConfigFilePath(), &config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := strings.Replace(config.AuthenticationFilePath, "~", user.HomeDir, -1)
	_, err = toml.DecodeFile(p, &auth)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := cli.NewApp()
	app.Name = "twitter-cli"
	app.Usage = "twitter cli"
	app.Description = "command-line twitter client"
	app.Version = version

	app.Commands = []cli.Command{
		{
			Name:   "timeline",
			Usage:  "show timeline",
			Action: cmdTimeline,
		},
		{
			Name:   "search",
			Usage:  "search tweets",
			Action: cmdSearch,
		},
	}

	app.Action = cmdTimeline

	app.RunAndExitOnError()
}
