package guest

import G "final-project/entities/guest"

type Guest interface {
	Create(newGuest G.Guest) (G.Guest, error)
	GetAll() ([]G.Guest, error)
	GetByUserID(ID uint) ([]G.Guest, error)
	Update(guestUpdate G.Guest) (G.Guest, error)
	Delete(ID, userID uint) error
}
