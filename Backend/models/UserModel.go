package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Role string 

const (
	AdminRole Role = "Admin"
	UserRole Role = "User"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty"`
	Email    string             `json:"email,omitempty"`
	Password string             `json:"password,omitempty"`
	Role     Role               `json:"role,omitempty"`  //this make sures that The field Role can Accept only 2 values admin/user
}

/*
{
	"name":"Vishnu Menon",
	"email":"vishnu.jio@gmail.com",
	"password":"thissucks",
	"role":"Admin"
}
*/