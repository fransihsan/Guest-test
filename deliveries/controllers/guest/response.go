package guest

import (
	G "final-project/entities/guest"
	"time"
)

type ResponseCreateGuest struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	NoHP  string `json:"no_hp"`
	Date  time.Time
	Pesan string `json:"pesan"`
}

func ToResponseCreateGuest(Guest G.Guest) ResponseCreateGuest {
	return ResponseCreateGuest{
		ID:   Guest.ID,
		Name: Guest.Name,
		NoHP: Guest.NoHP,
		Date: Guest.Date,
		Pesan: Guest.Pesan,
	}
}

type ResponseGetGuests struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	NoHP  string `json:"no_hp"`
	Date  time.Time
	Pesan string `json:"pesan"`
}

func ToResponseGetGuests(Guests []G.Guest) []ResponseGetGuests {
	responses := make([]ResponseGetGuests, len(Guests))
	for i := 0; i < len(Guests); i++ {
		responses[i].ID = Guests[i].ID
		responses[i].Name = Guests[i].Name
		responses[i].NoHP = Guests[i].NoHP
		responses[i].Date = Guests[i].Date
		responses[i].Pesan = Guests[i].Pesan
	}
	return responses
}

type ResponseUpdateGuest struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	NoHP  string `json:"no_hp"`
	Date  time.Time
	Pesan string `json:"pesan"`
}

func ToResponseUpdateGuest(Guest G.Guest) ResponseUpdateGuest {
	return ResponseUpdateGuest{
		ID:   Guest.ID,
		Name: Guest.Name,
		NoHP: Guest.NoHP,
		Date: Guest.Date,
		Pesan: Guest.Pesan,
	}
}
