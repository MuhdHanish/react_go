package utils

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
    return c.Status(status).JSON(fiber.Map{
        "success": true,
        "message": message,
        "data":    data,
    })
}

func ErrorResponse(c *fiber.Ctx, status int, message string, err error) error {
    response := fiber.Map{
        "success": false,
        "message": message,
    }
    
    if err != nil {
        response["error"] = err.Error()
    }
    
    return c.Status(status).JSON(response)
}
