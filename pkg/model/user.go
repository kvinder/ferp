package model

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// User type
type User struct {
	ID         int
	EmployeeID string
	Username   string
	Password   string
	Name       string
	Sex        string
	Department string
	Email      string
	Telephone  string
	Roles      []string
	CreateDate string
	UpdateDate string
}

var userStorage []*User

//CreateUser database
func CreateUser(user *User) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	checkErr(err)
	sqlQuery := `INSERT INTO APP_USER (employee_id,username,password,name,sex,
	department,email,telephone,createDate,updateDate) VALUES (?,?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlQuery, user.EmployeeID, user.Username, hashedPassword, user.Name, user.Sex, user.Department, user.Email, user.Telephone, user.CreateDate, user.UpdateDate)
	checkErr(err)
	sqlQuery = `SELECT ID FROM APP_USER WHERE Employee_ID = ?`
	rows, err := db.Query(sqlQuery, user.EmployeeID)
	checkErr(err)
	var appUserID int
	for rows.Next() {
		err = rows.Scan(&appUserID)
	}
	sqlQuery = `INSERT INTO APP_ROLE (role_name, createDate, updateDate, app_user) 
	VALUES (?, ?, ?, ?)`
	for _, role := range user.Roles {
		_, err := db.Exec(sqlQuery, role, user.CreateDate, user.UpdateDate, appUserID)
		checkErr(err)
	}
}

//UpdateUser database
func UpdateUser(user *User) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	checkErr(err)
	sqlQuery := `UPDATE APP_USER SET employee_id = ?, username = ?, password = ?, name = ?, sex = ?,
	department = ?, email = ?, telephone = ?, updateDate = ? WHERE id = ?`
	_, err = db.Exec(sqlQuery, user.EmployeeID, user.Username, hashedPassword, user.Name, user.Sex, user.Department, user.Email, user.Telephone, user.UpdateDate, user.ID)
	checkErr(err)
	sqlQuery = `DELETE FROM APP_ROLE WHERE app_user = ?`
	_, err = db.Exec(sqlQuery, user.ID)
	checkErr(err)
	sqlQuery = `INSERT INTO APP_ROLE (role_name, createDate, updateDate, app_user) 
	VALUES (?, ?, ?, ?)`
	for _, role := range user.Roles {
		_, err := db.Exec(sqlQuery, role, user.UpdateDate, user.UpdateDate, user.ID)
		checkErr(err)
	}
}

//ListUsers database
func ListUsers() []User {
	rows, err := db.Query("SELECT Employee_ID, Username, Name, Department, createDate FROM APP_USER")
	checkErr(err)
	var allUser []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.EmployeeID, &user.Username, &user.Name, &user.Department, &user.CreateDate)
		checkErr(err)
		allUser = append(allUser, user)
	}
	return allUser
}

//GetUser database
func GetUser(employeeID string) User {
	sqlQuery := `SELECT ID,employee_id,username,password,name,sex,department,email,telephone,createDate,updateDate FROM APP_USER WHERE Employee_ID = ?`
	rowsAppUser, err := db.Query(sqlQuery, employeeID)
	checkErr(err)
	user := User{}
	for rowsAppUser.Next() {
		err = rowsAppUser.Scan(&user.ID, &user.EmployeeID, &user.Username, &user.Password, &user.Name, &user.Sex, &user.Department, &user.Email, &user.Telephone, &user.CreateDate, &user.UpdateDate)
		checkErr(err)
	}
	sqlQuery = `SELECT role_name FROM APP_ROLE WHERE app_user = ?`
	rowsAppRole, err := db.Query(sqlQuery, user.ID)
	checkErr(err)
	for rowsAppRole.Next() {
		var roleName string
		err = rowsAppRole.Scan(&roleName)
		checkErr(err)
		user.Roles = append(user.Roles, roleName)
	}
	return user
}

//GetByUsername database
func GetByUsername(username string) User {
	sqlQuery := `SELECT ID,employee_id,username,password,name,sex,department,email,telephone,createDate,updateDate FROM APP_USER WHERE username = ?`
	rowsAppUser, err := db.Query(sqlQuery, username)
	checkErr(err)
	user := User{}
	for rowsAppUser.Next() {
		err = rowsAppUser.Scan(&user.ID, &user.EmployeeID, &user.Username, &user.Password, &user.Name, &user.Sex, &user.Department, &user.Email, &user.Telephone, &user.CreateDate, &user.UpdateDate)
		checkErr(err)
	}
	sqlQuery = `SELECT role_name FROM APP_ROLE WHERE app_user = ?`
	rowsAppRole, err := db.Query(sqlQuery, user.ID)
	checkErr(err)
	for rowsAppRole.Next() {
		var roleName string
		err = rowsAppRole.Scan(&roleName)
		checkErr(err)
		user.Roles = append(user.Roles, roleName)
	}
	return user
}

//DeleteUser database
func DeleteUser(employeeID string) {
	for i, user := range userStorage {
		if user.EmployeeID == employeeID {
			userStorage = append(userStorage[:i], userStorage[i+1:]...)
		}
	}
}

//UserOnLogin server
func UserOnLogin(r *http.Request) (User, error) {
	autenUsername := r.URL.Query().Get("username")
	var userLogin User
	if len(autenUsername) <= 0 {
		return userLogin, fmt.Errorf("No data in parameter")
	}
	userLogin = GetByUsername(autenUsername)
	return userLogin, nil
}
