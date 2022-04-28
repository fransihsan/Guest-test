package main

import (
	"final-project/configs"
	_AuthController "final-project/deliveries/controllers/auth"
	_UserController "final-project/deliveries/controllers/user"
	_GuestController "final-project/deliveries/controllers/guest"
	"final-project/deliveries/routes"
	_AuthRepo "final-project/repositories/auth"
	_UserRepo "final-project/repositories/user"
	_GuestRepo "final-project/repositories/guest"
	"final-project/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig(false)
	db := utils.InitDB(config)

	authRepo := _AuthRepo.NewAuthRepository(db)
	userRepo := _UserRepo.NewUserRepository(db)
	guestRepo := _GuestRepo.NewGuestRepository(db)

	ac := _AuthController.NewAuthController(authRepo)
	uc := _UserController.NewUserController(userRepo)
	gc := _GuestController.NewGuestController(guestRepo)

	e := echo.New()

	routes.RegisterPaths(e, ac, uc, gc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
