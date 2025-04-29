package main

import "github.com/joaolima7/api1_goexpert/configs"

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic("could not load config")
	}
	println(config.DBDriver)
}
