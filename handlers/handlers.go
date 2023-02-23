package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ApiHandler struct {
	Method string
	Path string
	Handler func(ctx *fiber.Ctx, db *gorm.DB) error
}


var ApiHandlers []ApiHandler = []ApiHandler{
	{
		Method: "GET",
		Path: "/",
		Handler: func(ctx *fiber.Ctx, db *gorm.DB) error {
			return ctx.SendString("Hello World")
		},
	},
	{
		Method: "POST",
		Path: "/profile",
		Handler: CreateProfile,
	},
}


func GenerateHandlers(api *fiber.Router, db *gorm.DB) {
	for _, val := range ApiHandlers {
		handler := val
		switch handler.Method {
		case "GET":
			(*api).Get(handler.Path, func(ctx *fiber.Ctx)error{
				return handler.Handler(ctx, db)
			})
		case "POST":
			(*api).Post(handler.Path, func (ctx *fiber.Ctx) error {
				return handler.Handler(ctx, db)
			})
		case "PUT":
			(*api).Put(handler.Path, func (ctx *fiber.Ctx) error {
				return handler.Handler(ctx, db)
			})
		case "DELETE":
			(*api).Delete(handler.Path, func (ctx *fiber.Ctx) error {
				return handler.Handler(ctx, db)
			})
		case "PATCH":
			(*api).Patch(handler.Path, func (ctx *fiber.Ctx) error {
				return handler.Handler(ctx, db)
			})
		case "OPTIONS":
			(*api).Options(handler.Path, func (ctx *fiber.Ctx) error {
				return handler.Handler(ctx, db)
			})
		case "HEAD":
			(*api).Head(handler.Path, func (ctx *fiber.Ctx) error {
				return handler.Handler(ctx, db)
			}) 
		case "CONNECT":
			(*api).Connect(handler.Path, func (ctx *fiber.Ctx) error {
				return handler.Handler(ctx, db)
			})
		case "TRACE":
			(*api).Trace(handler.Path, func (ctx *fiber.Ctx) error {
				return handler.Handler(ctx, db)
			})
		default:
			(*api).Get(handler.Path, func (ctx *fiber.Ctx) error {
				return handler.Handler(ctx, db)
			})
		}
	}
}