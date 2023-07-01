package url

import (
	"github.com/DaffaAudyaPramana/proyeks-2/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage)
	page.Get("/user", controller.GetUser)               //API from user whatsapp message from iteung gowa
	page.Get("/pendaftaran", controller.GetPendaftaran) //API from user whatsapp message from iteung gowa
	page.Get("/pembayaran", controller.GetPembayaran)   //API from user whatsapp message from iteung gowa
	page.Get("/pengumuman", controller.GetPengumuman)   //API from user whatsapp message from iteung gowa
	page.Get("/kursus", controller.GetKursus)           //API from user whatsapp message from iteung gowa

	// page.Get("/", controller.Sink)
	// page.Post("/", controller.Sink)
	// page.Put("/", controller.Sink)
	// page.Patch("/", controller.Sink)
	// page.Delete("/", controller.Sink)
	// page.Options("/", controller.Sink)

}
