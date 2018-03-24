package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

var (
	version  string
	revision string
)

var config globalConfig

const (
	configFileName = "config.toml"
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

func mainAction(c *cli.Context) error {

	return nil
}

func main() {

	configFile := getConfigFilePath()
	_, err := os.Stat(configFile)
	if err != nil {

		config = getDefaultConfig()
		err = config.save(configFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	} else {

		_, err := toml.DecodeFile(getConfigFilePath(), &config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	app := cli.NewApp()
	app.Name = "twitter-cli"
	app.Usage = "twitter cli"
	app.Description = "command-line twitter client"
	app.Version = version

	app.Action = mainAction

	app.RunAndExitOnError()
}
