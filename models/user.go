package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Role     string
}

type Permission struct {
	ID     uint `gorm:"primaryKey"`
	Role   string
	Module string
	Access string
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"     binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) RegisterUser(db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return db.Create(u).Error
}

func (u *User) AuthenticateUser(db *gorm.DB) error {
	var user User
	if err := db.Where("username = ?", u.Username).First(&user).Error; err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		return err
	}
	u.ID = user.ID
	u.Role = user.Role
	return nil
}

func (p *Permission) CheckPermission(db *gorm.DB) bool {
	if err := db.Where("role = ? AND module = ? AND access = ?", p.Role, p.Module, p.Access).First(p).Error; err != nil {
		return false
	}
	return true
}
