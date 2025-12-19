package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"size:50;uniqueIndex;not null" json:"username"` // 用户名50字符足够
	Email     string    `gorm:"size:255;uniqueIndex;not null" json:"email"`   // 邮箱标准最大255
	Password  string    `gorm:"size:255;not null" json:"-"`                   // 密码哈希通常60-255字符
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`             // 自动填充时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`             // 自动更新时间
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
