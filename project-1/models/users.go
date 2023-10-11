package models

type Users struct{
	Name string
	Email string
	Password string
	Permissions map[string]bool // no need to make()
	Admin bool
}

