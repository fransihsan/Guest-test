package guest

import (
	G "final-project/entities/guest"

	"gorm.io/gorm"
)

type RequestCreateGuest struct {
	Name  string `json:"name" form:"name"`
	NoHP  string `json:"no_hp" form:"no_hp"`
	Pesan string `json:"pesan" form:"pesan"`
}

func (Req RequestCreateGuest) ToEntityGuest(UserID uint) G.Guest {
	return G.Guest{
		Name: Req.Name,
		NoHP: Req.NoHP,
		Pesan: Req.Pesan,
		UserID: UserID,
	}
}

type RequestUpdateGuest struct {
	Name     string `json:"name" form:"name"`
	NoHP  string `json:"no_hp" form:"no_hp"`
	Pesan string `json:"pesan" form:"pesan"`
}

func (Req RequestUpdateGuest) ToEntityGuest(ID, UserID uint) G.Guest {
	return G.Guest{
		Model: gorm.Model{ID: ID},
		Name:  Req.Name,
		NoHP: Req.NoHP,
		Pesan: Req.Pesan,
		UserID: UserID,
	}
}
