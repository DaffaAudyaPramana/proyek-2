package controller

import (
	"github.com/DaffaAudyaPramana/pkbackend"
	"github.com/DaffaAudyaPramana/proyek-2/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/whatsauth"
)

var DataUser = "user"
var DataPendaftaran = "pendaftaran"
var DataPembayaran = "pembayaran"
var DataPengumuman = "pengumuman"
var DataKursus = "kursus"

// type HTTPRequest struct {
// 	Header string `json:"header"`
// 	Body   string `json:"body"`
// }

// func Sink(c *fiber.Ctx) error {
// 	var req HTTPRequest
// 	req.Header = string(c.Request().Header.Header())
// 	req.Body = string(c.Request().Body())
// 	return c.JSON(req)
// }

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

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetUser(c *fiber.Ctx) error {
	getstatus := pkbackend.GetDataUser("Arya")
	return c.JSON(getstatus)
}

func GetPendaftaran(c *fiber.Ctx) error {
	getstatus := pkbackend.GetDataPendaftaran("92348927348")
	return c.JSON(getstatus)
}

func GetPembayaran(c *fiber.Ctx) error {
	getstatus := pkbackend.GetDataPembayaran("Dibayar?")
	return c.JSON(getstatus)
}

func GetPengumuman(c *fiber.Ctx) error {
	getstatus := pkbackend.GetDataPengumuman("Fighting")
	return c.JSON(getstatus)
}
func GetKursus(c *fiber.Ctx) error {
	getstatus := pkbackend.GetDataKursus("Jon Snow")
	return c.JSON(getstatus)
}
