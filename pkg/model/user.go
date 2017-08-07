package model

import "time"

// User type
type User struct {
	EmployeeID string
	Username   string
	Password   string
	Name       string
	Sex        string
	Department string
	Email      string
	Telephone  string
	Role       []string
	CreateDate time.Time
	UpdateDate time.Time
}

var userStorage []*User

//CreateUser database
func CreateUser(user *User) {
	userStorage = append(userStorage, user)
}

//ListUsers database
func ListUsers() []*User {
	return userStorage
}

//GetUser database
func GetUser(employeeID string) *User {
	for _, user := range userStorage {
		if user.EmployeeID == employeeID {
			return user
		}
	}
	return nil
}

//DeleteUser database
func DeleteUser(employeeID string) {
	for i, user := range userStorage {
		if user.EmployeeID == employeeID {
			userStorage = append(userStorage[:i], userStorage[i+1:]...)
		}
	}
}
