package main

import "github.com/oaraujocesar/go-expert-api/configs"

func main() {
	// TODO: implement

	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)

}
