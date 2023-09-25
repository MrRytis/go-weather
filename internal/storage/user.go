package storage

import "time"

type User struct {
	Id       int32
	Email    string
	Password string
	CreateAt time.Time
	UpdateAt time.Time
}

func (r *Repository) CreateUser(u User) error {
	return r.Db.Create(&u).Error
}

func (r *Repository) GetUserByEmail(email string) (User, error) {
	var u User
	err := r.Db.Where("email = ?", email).First(&u).Error

	return u, err
}
