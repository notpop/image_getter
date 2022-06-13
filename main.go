package main

import (
	"fmt"
	"image_getter/config"
	"io"
	"net/http"
	"os"
)

const (
	IMAGE_DIRECTORY_PATH = "images/test/"
	IMAGE_PATH           = "save.jpeg"
)

func main() {
	response, err := http.Get(config.Config.TargetUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	err = os.MkdirAll(IMAGE_DIRECTORY_PATH, 0777)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create(IMAGE_DIRECTORY_PATH + IMAGE_PATH)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
}
