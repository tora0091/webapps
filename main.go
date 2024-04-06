package main

import (
	"github.com/joho/godotenv"
	"github.com/tora0091/webapps/environments"
	"github.com/tora0091/webapps/routers"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error don't open .env file.")
	}
}

func main() {
	env := environments.NewEnv()
	routers.NewRouters(env).RouterUp()
}
