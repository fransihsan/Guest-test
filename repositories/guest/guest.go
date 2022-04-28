package guest

import (
	"errors"
	G "final-project/entities/guest"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type GuestRepository struct {
	db *gorm.DB
}

func NewGuestRepository(db *gorm.DB) *GuestRepository {
	return &GuestRepository{
		db: db,
	}
}

func (repo *GuestRepository) Create(newGuest G.Guest) (G.Guest, error) {
	if err := repo.db.Create(&newGuest).Error; err != nil {
		log.Warn(err)
		return G.Guest{}, errors.New("gagal membuat daftar tamu baru")
	}
	return newGuest, nil
}

func (repo *GuestRepository) GetAll() ([]G.Guest, error) {
	guests := []G.Guest{}
	repo.db.Find(&guests)
	if len(guests) < 1 {
		return nil, errors.New("tidak terdapat daftar tamu sama sekali")
	}
	return guests, nil
}

func (repo *GuestRepository) GetByUserID(ID uint) ([]G.Guest, error) {
	guests := []G.Guest{}
	repo.db.Find(&guests, ID)
	if len(guests) < 1 {
		return nil, errors.New("tidak terdapat daftar tamu sama sekali")
	}
	return guests, nil
}

func (repo *GuestRepository) Update(guestUpdate G.Guest) (G.Guest, error) {
	if rowsAffected := repo.db.Model(&guestUpdate).Updates(guestUpdate).RowsAffected; rowsAffected == 0 {
		return G.Guest{}, errors.New("gagal memperbaharui daftar tamu")
	}

	repo.db.First(&guestUpdate)
	return guestUpdate, nil
}

func (repo *GuestRepository) Delete(ID, userID uint) error {
	if rowsAffected := repo.db.Delete(&G.Guest{}, ID, userID).RowsAffected; rowsAffected == 0 {
		return errors.New("tidak ada daftar tamu yang dihapus")
	}
	return nil
}

