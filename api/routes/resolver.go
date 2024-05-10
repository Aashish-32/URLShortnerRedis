package routes

import (
	"github.com/Aashish-32/URL-Shortener/database"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.Close()
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in the database",
			"value": err,
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to db",
			"value": err,
		})

	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301) //301=permanent redirection.302=temporary

}
