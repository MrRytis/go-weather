package storage

import "time"

type User struct {
	Id        int32     `gorm:"primaryKey"`
	Email     string    `gorm:"unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (r *Repository) CreateUser(u User) error {
	return r.Db.Create(&u).Error
}

func (r *Repository) GetUserByEmail(email string) (User, error) {
	var u User
	err := r.Db.Where("email = ?", email).First(&u).Error

	return u, err
}

func (r *Repository) UpdateUser(user *User) *User {
	r.Db.Save(&user)

	return user
}
