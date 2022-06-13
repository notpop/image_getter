package main

import (
	"image_getter/config"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get(config.Config.TargetUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	file, err := os.Create("save.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
}
