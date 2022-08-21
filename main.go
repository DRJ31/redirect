package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

type URLInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func readData() map[string]string {
	var urlInfos []URLInfo
	urlMap := make(map[string]string)

	jsonFile, err := os.Open("urls.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&urlInfos)
	if err != nil {
		log.Fatal(err)
	}

	for _, urlInfo := range urlInfos {
		urlMap[urlInfo.Url] = urlInfo.Name
	}

	return urlMap
}

func redirectController(ctx *fiber.Ctx) error {
	urlMap := readData()
	dst := urlMap[ctx.Hostname()]

	if len(dst) == 0 {
		return ctx.Redirect("https://uichcc.com")
	} else {
		return ctx.Redirect(fmt.Sprintf("https://uichcc.com/%v/", dst))
	}
}

func main() {
	app := fiber.New()
	app.Get("/", redirectController)
	app.Listen("0.0.0.0:5000")
}
