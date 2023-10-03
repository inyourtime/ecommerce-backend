package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type providerType string
type roleType string

const (
	LocalProvider  providerType = "local"
	GoogleProvider providerType = "google"
)

const (
	AdminRole roleType = "admin"
	UserRole  roleType = "user"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Provider  providerType       `json:"provider,omitempty" bson:"provider"`
	Email     string             `json:"email,omitempty" bson:"email"`
	Password  string             `json:"password,omitempty" bson:"password"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname"`
	Avatar    string             `json:"avater,omitempty" bson:"avater,omitempty"`
	Role      roleType           `json:"role,omitempty" bson:"role"`
	GoogleID  string             `json:"googleID,omitempty" bson:"googleId,omitempty"`
	IsActive  bool               `json:"isActive,omitempty" bson:"isActive"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
}
