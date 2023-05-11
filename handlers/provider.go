package handlers

import (
	"clicker/middlewares"
	"clicker/services"

	"github.com/gofiber/fiber/v2"
)

func SetUpHandlers(app *fiber.App, cs *services.ClickerService) {
	api := app.Group("/api")
	authRoutes(api, cs)
	postRoutes(api, cs)
	portalRoutes(api, cs)
}

func authRoutes(router fiber.Router, service *services.ClickerService) {
	auth := router.Group("auth")
	ah := NewAuthHandler(*service)
	auth.Post("/signup", ah.Register)
	auth.Post("/login", ah.Login)
}

func portalRoutes(router fiber.Router, service *services.ClickerService) {
	portal := router.Group("portal")
	ph := NewPortalHandler(*service)
	portal.Get("findAllUser", ph.FindAllUser)
	portal.Post("createPost", ph.CreatePost)
	portal.Delete("deletePost", ph.DeletePost)
	portal.Get("post/:id", ph.FindPostDetail)
	portal.Get("login", ph.Login)
}

func postRoutes(router fiber.Router, service *services.ClickerService) {
	post := router.Group("posts")
	poH := NewPostHandler(*service)
	post.Get("/", poH.GetAllPost)
	post.Get("/:id", poH.GetPostById)
	post.Use(middlewares.JwtGuard())
	post.Post("/markAsClicked", poH.MarkAsClicked)
}
