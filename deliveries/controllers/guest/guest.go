package guest

import (
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	_GuestRepo "final-project/repositories/guest"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GuestController struct {
	repo _GuestRepo.Guest
}

func NewGuestController(repository _GuestRepo.Guest) *GuestController {
	return &GuestController{
		repo: repository,
	}
}

func (ctl *GuestController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)
		NewGuest := RequestCreateGuest{}

		if err := c.Bind(&NewGuest); err != nil || NewGuest.Name == "" || NewGuest.Pesan == "" || NewGuest.NoHP == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, nama, no hp dan pesan tamu tidak boleh kosong"))
		}

		res, err := ctl.repo.Create(NewGuest.ToEntityGuest(uint(userID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan daftar tamu baru", ToResponseCreateGuest(res)))
	}
}

func (ctl *GuestController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed jwt"))
		}
		res, err := ctl.repo.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua daftar tamu", res))
	}
}

func (ctl *GuestController) GetByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)
		res, err := ctl.repo.GetByUserID(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua daftar tamu berdasarkan user id", res))
	}
}

func (ctl *GuestController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)
		ID, _ := strconv.Atoi(c.Param("id"))
		var updateGuest RequestUpdateGuest

		if err := c.Bind(&updateGuest); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("terdapat kesalahan input dari user"))
		}

		res, err := ctl.repo.Update(updateGuest.ToEntityGuest(uint(ID), uint(userID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses memperbarui daftar tamu", ToResponseUpdateGuest(res)))
	}


}

func (ctl *GuestController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)
		ID, _ := strconv.Atoi(c.Param("id"))

		err := ctl.repo.Delete(uint(ID), uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus user", err))
	}
}

