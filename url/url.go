package url

import (
	"github.com/DaffaAudyaPramana/proyek-2/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage)
	page.Get("/user", controller.Getuser) //API from user whatsapp message from iteung gowa
	// page.Get("/pendaftaran", controller.Getpendaftaran)         //API from user whatsapp message from iteung gowa
	// page.Get("/pembayaran", controller.Getpembayaran)   //API from user whatsapp message from iteung gowa
	// page.Get("/pengumuman", controller.Getpengumuman) //API from user whatsapp message from iteung gowa
	// page.Get("/kursus", controller.Getkursus) //API from user whatsapp message from iteung gowa

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

}
