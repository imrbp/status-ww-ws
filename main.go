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
	go helper.ReadJson(filePath, &data)

	var statusWind, statusWater string
	app := fiber.New()

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
				statusWind = CheckWind(data.Status.Wind)
				statusWater = CheckWater(data.Status.Water)

				fmt.Println("\nJSON READ ")
				fmt.Printf("Water :%d, %s \n", data.Status.Water, statusWater)
				fmt.Printf("Wind :%d, %s \n", data.Status.Wind, statusWind)
				err := c.WriteJSON(data)
				if err != nil {
					return
				}
			case <-fileChanges.C:
				go helper.WriteJson(filePath, &data)
				fmt.Println("\nFile Changes \ntimes :", <-fileRead.C)
				err := c.WriteJSON("JSON UPDATED")
				if err != nil {
					return
				}
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
}

func CheckWind(value int) string {
	switch {
	case value < 5:
		return "Aman"
	case value >= 6 && value <= 8:
		return "Siaga"
	default:
		return "Bahaya"
	}
}

func CheckWater(value int) string {
	switch {
	case value > 15:
		return "Bahaya"
	case value >= 6:
		return "Siaga"
	default:
		return "Aman"

	}
}
