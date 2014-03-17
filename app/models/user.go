package models

type User struct {
	Id       uint   // database generated
	Email    string // verified email address for logon authentication and contact
	Password string // should be a salted hash string
}
