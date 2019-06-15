package main

type globalConfig struct {
	AuthenticationFilePath string `toml:"authentication_file_path"`
}

type authentication struct {
	AccessToken       string `toml:"access_token"`
	AccessTokenSecret string `toml:"access_token_secret"`
	ConsumerKey       string `toml:"consumer_key"`
	ConsumerSecret    string `toml:"consumer_secret"`
}
