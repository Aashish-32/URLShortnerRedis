package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/Aashish-32/URL-Shortener/database"
	"github.com/Aashish-32/URL-Shortener/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type request struct {
	URL         string        `json:"url"`
	Expire      time.Duration `json:"expire"`
	CustomShort string        `json:"custom_short"`
}

type response struct {
	URL            string        `json:"url"`
	Short          string        `json:"short"`
	Expire         time.Duration `json:"expire"`
	RateLimiting   int           `json:"rate_limiting"`
	Ratelimitreset time.Duration `json:"ratelimitreset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "cannot parse json", "error": err})

	}

	//implement rate limiting....{//*check ip of the user calling and if it is in the db or not.if it is decrement by 1}
	r2 := database.CreateClient(1)
	defer r2.Close()

	_, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		value, _ := r2.Get(database.Ctx, c.IP()).Result()
		valint, _ := strconv.Atoi(value)

		if valint <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "rate limit exceeded",
				"limit": limit / time.Minute,
			})
		}

	}

	//check if input is an actual URL

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid url"})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "you cant exploit it"})

	}

	//enforcing https,ssl
	body.URL = helpers.EnforceHTTPS(body.URL)
	r2.Decr(database.Ctx, c.IP())

	return nil

}
