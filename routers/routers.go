package routers

import (
	"github.com/gofiber/fiber/v2"
)

func getmeta(route string) string {
	MetaTitle := "Hello, Meta!"
	if route == "/" {
		MetaTitle = "Home"
	} else if route == "features" {
		MetaTitle = "Features"
	} else if route == "pricing" {
		MetaTitle = "Pricing"
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

	router.Get("/features", func(c *fiber.Ctx) error {
		MetaTitle := getmeta("features")
		return c.Render("features", fiber.Map{
			"MetaTitle": MetaTitle,
		})
	})

	router.Get("/pricing", func(c *fiber.Ctx) error {
		MetaTitle := getmeta("pricing")
		return c.Render("pricing", fiber.Map{
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
		//  header.component.html = header.component
		return c.Render("components/headers/header.component", fiber.Map{
			"Title": "Hello, Header!",
		})
	})

	components.Get("/footer", func(c *fiber.Ctx) error {
		//  footer.component.html = footer.component
		return c.Render("components/footers/footer.component", fiber.Map{
			"Title": "Hello, Footer!",
		})
	})

	components.Get("/home", func(c *fiber.Ctx) error {
		// home.component.html = home.component
		return c.Render("components/home.component", fiber.Map{})
	})

	components.Get("/features", func(c *fiber.Ctx) error {

		// features.component.html = features.component
		return c.Render("components/features.component", fiber.Map{})
	})
	components.Get("/pricing", func(c *fiber.Ctx) error {

		// pricing.component.html = pricing.component
		return c.Render("components/pricing.component", fiber.Map{})
	})
}
