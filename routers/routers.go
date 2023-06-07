package routers

import (
	"github.com/gofiber/fiber/v2"
)

func getmeta(route string) string {
	MetaTitle := "Hello, Meta!"
	if route == "/" {
		MetaTitle = "Home"
	} else if route == "about" {
		MetaTitle = "About"
	}
	return MetaTitle
}

// Capitalize the first letter of the function name to export it
func RegisterRoutes(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		MetaTitle := getmeta("/")
		return c.Render("index", fiber.Map{
			"Title":     "Hello, World!",
			"MetaTitle": MetaTitle,
		})
	})

	router.Get("/about", func(c *fiber.Ctx) error {
		MetaTitle := getmeta("about")
		return c.Render("about", fiber.Map{
			"MetaTitle": MetaTitle,
		})
	})

	// component groups

	components := router.Group("/components")

	components.Get("/meta", func(c *fiber.Ctx) error {
		route := c.Query("route")
		MetaTitle := getmeta(route)
		return c.Render("partials/meta", fiber.Map{
			"MetaTitle": MetaTitle,
		})
	})

	components.Get("/header", func(c *fiber.Ctx) error {
		return c.Render("components/headers/header", fiber.Map{
			"Title": "Hello, Header!",
		})
	})

	components.Get("/footer", func(c *fiber.Ctx) error {
		return c.Render("components/footers/footer", fiber.Map{
			"Title": "Hello, Footer!",
		})
	})

	components.Get("/home", func(c *fiber.Ctx) error {
		return c.Render("components/home", fiber.Map{})
	})

	components.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("components/about", fiber.Map{})
	})
}
