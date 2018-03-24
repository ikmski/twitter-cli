package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type globalConfig struct {
	AccessToken       string `toml:"access_token"`
	AccessTokenSecret string `toml:"access_token_secret"`
	ConsumerKey       string `toml:"consumer_key"`
	ConsumerSecret    string `toml:"consumer_secret"`
}

func getDefaultConfig() globalConfig {
	return globalConfig{
		AccessToken:       "",
		AccessTokenSecret: "",
		ConsumerKey:       "",
		ConsumerSecret:    "",
	}
}

func (c *globalConfig) save(file string) error {

	dir := filepath.Dir(file)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(c)
	if err != nil {
		return err
	}

	return nil
}
