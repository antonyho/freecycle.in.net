package models

type User struct {
	Email    string // verified email address for logon authentication and contact
	Password string // should be a salted hash string
}
