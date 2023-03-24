package controller

import (
	"github.com/DaffaAudyaPramana/pkbackend"
	"github.com/DaffaAudyaPramana/proyek-2/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/whatsauth"
)

// var DataUser = "Daffa Audya Pramana"

type HTTPRequest struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

func Sink(c *fiber.Ctx) error {
	var req HTTPRequest
	req.Header = string(c.Request().Header.Header())
	req.Body = string(c.Request().Body())
	return c.JSON(req)
}

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func Getdbuser(c *fiber.Ctx) error {
	getstatus := pkbackend.GetDataUser("Daffa Audya Pramana")
	return c.JSON(getstatus)
}

// func Getpembayaran(c *fiber.Ctx) error {
// 	getstatus := backend.GetDataPembayaran("Apakah kurikulum ini menyulitkan hidup?")
// 	return c.JSON(getstatus)
// }
// func Getpendaftaran(c *fiber.Ctx) error {
// 	getstatus := backend.GetDataPendaftaran("0822126722")
// 	return c.JSON(getstatus)
// }
// func Getpengumuman(c *fiber.Ctx) error {
// 	getstatus := backend.GetDataPengumuman("Denmark")
// 	return c.JSON(getstatus)
// }
// func Getprogramkursus(c *fiber.Ctx) error {
// 	getstatus := backend.GetDataProgramKursus("johndoe@flex.co")
// 	return c.JSON(getstatus)
// }

// func GetHome(c *fiber.Ctx) error {
// 	getip := backend
// 	return c.JSON(getip)
