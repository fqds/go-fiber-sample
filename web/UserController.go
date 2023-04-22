package web

import (
	"go-fiber-sample/dto"
	"go-fiber-sample/service"
	request2 "go-fiber-sample/web/request"
	response2 "go-fiber-sample/web/response"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			request2.UserRequest
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		log.Println(req)

		user := &dto.UserDto{
			Name:     req.Name,
			Password: req.Password,
		}
		
		if err := service.CreateUser(user); err != nil {
			return c.Status(401).JSON(err.Error())
		}

		signedToken, err := service.CreateSession(user)
		response := &response2.AuthTokenResponse{AuthToken: signedToken}
		if err != nil {
			return c.Status(401).JSON(err.Error())
		}
		return c.Status(201).JSON(response)
	}
}

func CreateSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			request2.UserRequest
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		log.Println(req)

		user := &dto.UserDto{
			Name:     req.Name,
			Password: req.Password,
		}
		signedToken, err := service.CreateSession(user)
		response := &response2.AuthTokenResponse{AuthToken: signedToken}
		if err != nil {
			return c.Status(401).JSON(err.Error())
		}
		return c.Status(200).JSON(response)
	}
}

func ApproveSession() fiber.Handler {
	return func(c * fiber.Ctx) error {
		type request struct {
			request2.AuthTokenRequest
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		log.Println(req)
		
		user, err := service.ApproveSession(req.AuthToken)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		c.Locals("user", user)
		return c.Status(200).Next()
	}
}