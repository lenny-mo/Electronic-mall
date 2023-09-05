package main

import (
	"eletronicMall/config"
	"eletronicMall/routes"
)

func main() {
	println("Hello, World!")

	config.Init()

	r := routes.NewRouter()
	err := r.Run(config.HttpPort)
	if err != nil {
		panic(err)
	}
        
}
