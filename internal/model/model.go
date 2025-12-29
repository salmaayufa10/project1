package model

import "time"

type Book struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Publisher       string    `json:"publisher"`
	Year            int       `json:"year"`
	Isbn            int       `json:"isbn"`
	TotalCopies     int       `json:"total_copies"`
	AvailableCopies int       `json:"available_copies"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type Books struct {
	Books []Book `json:"books"`
}

type Lib_user struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Lib_users struct {
	Lib_users []Lib_user `json:"lib_users"`
}
