package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleWare(ctx *fiber.Ctx) error {

	current_path := ctx.Path()

	switch current_path {
	case "/v1/" :
		return ctx.Next()
	default:
		return func (ctx *fiber.Ctx) error {
			// TODO: Implement auth middleware
			return ctx.Next()
		}(ctx)
	}

}