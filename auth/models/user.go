package models

import (
	"time"

	"github.com/rs401/myauth/pb"
	"gorm.io/gorm"
)

// User model for user
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  []byte `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) ToProtoBuffer() *pb.User {
	return &pb.User{
		Id:        uint64(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		Password:  string(u.Password),
		CreatedAt: u.CreatedAt.Unix(),
		UpdatedAt: u.UpdatedAt.Unix(),
	}
}

func (u *User) FromProtoBuffer(user *pb.User) {
	u.ID = uint(user.GetId())
	u.Name = user.GetName()
	u.Email = user.GetEmail()
	u.Password = []byte(user.GetPassword())
	u.CreatedAt = time.Unix(user.CreatedAt, 0)
	u.UpdatedAt = time.Unix(user.UpdatedAt, 0)
}
