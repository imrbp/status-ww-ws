package main

import (
	"fmt"
	"log"
	"microservice/entity"
	"microservice/helper"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/contrib/websocket"
)

func main() {

	fileChanges := time.NewTicker(15 * time.Second)
	fileRead := time.NewTicker(1 * time.Second)
	filePath := "./data.json"

	var data entity.Data

	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("client.html", fiber.Map{
			"Hello ": "World",
		})
	})
	app.Get("/ws/status", websocket.New(func(c *websocket.Conn) {
		for {
			select {
			case <-fileRead.C:
				go helper.ReadJson(filePath, &data)
				c.WriteJSON(data)
				fmt.Println("Water :", data.Status.Water)
				fmt.Println("Wind :", data.Status.Wind)
			case <-fileChanges.C:
				go helper.WriteJson(filePath, &data)
				fmt.Println("times :", <-fileRead.C)
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
}
