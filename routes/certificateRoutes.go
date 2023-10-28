package routes

import (
	certificatecontroller "github.com/Dandy-CP/go-fiber-portfolio/controllers/certificateController"
	"github.com/Dandy-CP/go-fiber-portfolio/middleware"
	"github.com/gofiber/fiber/v2"
)

func CertificateRoutes(certificate fiber.Router) {
	certificate.Get("/", certificatecontroller.GetListCertificate)

	certificate.Post("/",
		middleware.ValidateCertificate,
			certificatecontroller.CreateCertificate)

	certificate.Put("/:id",
		middleware.ValidateCertificate,
			certificatecontroller.UpdateCertificate)

	certificate.Delete("/:id", certificatecontroller.DeleteCertificate)
}