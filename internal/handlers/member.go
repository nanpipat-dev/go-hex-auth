package handlers

import (
	"go-hex-auth/internal/core/domain"
	"go-hex-auth/internal/core/services"

	"github.com/gofiber/fiber/v2"
)

type MemberHandlers struct {
	service services.MemberServiceInterface
}

func NewMemberHandlers(service services.MemberServiceInterface) *MemberHandlers {
	return &MemberHandlers{
		service: service,
	}
}

func (h *MemberHandlers) CreateMember(c *fiber.Ctx) error {
	var newMember domain.MembersRequest

	var err error

	err = c.BodyParser(&newMember)
	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON("Error incorrect input syntax")
	}

	err = h.service.CreateMember(newMember)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": &fiber.Map{
				"code": fiber.StatusBadRequest,
				"message": []string{
					err.Error(),
				},
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": &fiber.Map{
			"code": fiber.StatusOK,
			"message": []string{
				"Success",
			},
		},
	})
}

func (h *MemberHandlers) GetMember(c *fiber.Ctx) error {
	id := c.Query("id")
	members, err := h.service.GetMember(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": &fiber.Map{
				"code": fiber.StatusBadRequest,
				"message": []string{
					err.Error(),
				},
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": &fiber.Map{
			"code": fiber.StatusOK,
			"message": []string{
				"Success",
			},
		},
		"data": members,
	})
}

func (h *MemberHandlers) Login(c *fiber.Ctx) error {
	var request domain.MembersRequest
	var err error

	err = c.BodyParser(&request)
	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON("Error incorrect input syntax")
	}

	if request.Username == "" {
		return c.
			Status(fiber.StatusBadRequest).
			JSON("username or password invalid")
	}

	if request.Password == "" {
		return c.
			Status(fiber.StatusBadRequest).
			JSON("username or password invalid")
	}

	member, err := h.service.Login(request.Username, request.Password)
	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(&fiber.Map{
				"status": &fiber.Map{
					"code": fiber.StatusBadRequest,
					"message": []string{
						err.Error(),
					},
				},
			})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": &fiber.Map{
			"code": fiber.StatusOK,
			"message": []string{
				"Success",
			},
		},
		"data": member,
	})

}
