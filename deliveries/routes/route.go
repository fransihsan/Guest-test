package routes

import (
	"final-project/deliveries/controllers/auth"
	"final-project/deliveries/controllers/guest"
	"final-project/deliveries/controllers/user"
	"final-project/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaths(e *echo.Echo, ac *auth.AuthController, uc *user.UserController, gc *guest.GuestController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Auth Route
	a := e.Group("/login")
	a.POST("", ac.Login())

	// User Route
	u := e.Group("/users")
	u.POST("", uc.Create())
	u.GET("/list", uc.GetAllUsers())
	uj := u.Group("/jwt")
	uj.Use(middlewares.JWTMiddleware())
	uj.GET("/:id", uc.GetByID())
	uj.GET("/me", uc.Get())
	uj.PUT("/me", uc.Update())
	uj.DELETE("/:id", uc.Delete())

	// Guest Route
	g := e.Group("/guest")
	g.Use(middlewares.JWTMiddleware())
	g.GET("", gc.GetAll())
	g.POST("", gc.Create())
	g.GET("/me", gc.GetByUserID())
	g.PUT("/:id", gc.Update())
	g.DELETE("/:id", gc.Delete())
}
