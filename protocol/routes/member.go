package routes

import (
	"go-hex-auth/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func MemberRoutes(api fiber.Router, handler *handlers.MemberHandlers) {
	{
		api.Post("/create", handler.CreateMember)
		api.Get("/get", handler.GetMember)
		api.Post("/login", handler.Login)
		api.Post("/refresh", handler.Refresh)
	}
}
