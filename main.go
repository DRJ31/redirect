package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

type URLInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func loadURLMap() map[string]string {
	jsonFile, err := os.Open("urls.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	var urlInfos []URLInfo
	if err = json.NewDecoder(jsonFile).Decode(&urlInfos); err != nil {
		log.Fatal(err)
	}

	urlMap := make(map[string]string, len(urlInfos))
	for _, u := range urlInfos {
		urlMap[u.Url] = u.Name
	}
	return urlMap
}

func main() {
	urlMap := loadURLMap()

	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		dst := urlMap[c.Hostname()]
		if dst == "" {
			return c.Redirect().To("https://uichcc.app")
		}
		return c.Redirect().To(fmt.Sprintf("https://uichcc.app/%v/", dst))
	})

	log.Fatal(app.Listen("0.0.0.0:5000"))
}
