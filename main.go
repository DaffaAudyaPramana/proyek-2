package main

import (
	"log"

	"github.com/DaffaAudyaPramana/proyeks-2/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/DaffaAudyaPramana/proyeks-2/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
