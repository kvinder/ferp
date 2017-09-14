package model

import (
	"fmt"
	"net/http"
	"strconv"

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
	db := getConnection()
	defer db.Close()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	checkErr(err)
	sqlQuery := `INSERT INTO APP_USER (employee_id,username,password,name,sex,
	department,email,telephone,createDate,updateDate) VALUES (?,?,?,?,?,?,?,?,?,?)`
	result, err := db.Exec(sqlQuery, user.EmployeeID, user.Username, hashedPassword, user.Name, user.Sex, user.Department, user.Email, user.Telephone, user.CreateDate, user.UpdateDate)
	checkErr(err)
	appUserID, err := result.LastInsertId()
	checkErr(err)
	sqlQuery = `SELECT ID FROM APP_ROLE WHERE role_name = ?`
	for _, role := range user.Roles {
		rows, err := db.Query(sqlQuery, role)
		checkErr(err)
		var roleID string
		if rows.Next() {
			rows.Scan(&roleID)
		}
		rows.Close()
		i, _ := strconv.Atoi(roleID)
		checkErr(err)
		sqlInsertRoles := `INSERT INTO APP_USERS_APP_ROLES (app_user,app_role) VALUES (?,?)`
		_, err = db.Exec(sqlInsertRoles, int(appUserID), i)
		checkErr(err)
	}
}

//UpdateUser database
func UpdateUser(user *User) {
	db := getConnection()
	defer db.Close()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	checkErr(err)
	sqlQuery := `UPDATE APP_USER SET employee_id = ?, username = ?, password = ?, name = ?, sex = ?,
	department = ?, email = ?, telephone = ?, updateDate = ? WHERE id = ?`
	_, err = db.Exec(sqlQuery, user.EmployeeID, user.Username, hashedPassword, user.Name, user.Sex, user.Department, user.Email, user.Telephone, user.UpdateDate, user.ID)
	checkErr(err)
	sqlQuery = `DELETE FROM APP_USERS_APP_ROLES WHERE app_user = ?`
	_, err = db.Exec(sqlQuery, user.ID)
	checkErr(err)
	sqlQuery = `SELECT ID FROM APP_ROLE WHERE role_name = ?`
	for _, role := range user.Roles {
		rows, err := db.Query(sqlQuery, role)
		checkErr(err)
		var roleID string
		if rows.Next() {
			rows.Scan(&roleID)
		}
		rows.Close()
		i, _ := strconv.Atoi(roleID)
		checkErr(err)
		sqlInsertRoles := `INSERT INTO APP_USERS_APP_ROLES (app_user,app_role) VALUES (?,?)`
		_, err = db.Exec(sqlInsertRoles, user.ID, i)
		checkErr(err)
	}
}

//ListUsers database
func ListUsers() []User {
	db := getConnection()
	defer db.Close()
	rows, err := db.Query("SELECT Employee_ID, Username, Name, Department, createDate FROM APP_USER")
	checkErr(err)
	var allUser []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.EmployeeID, &user.Username, &user.Name, &user.Department, &user.CreateDate)
		checkErr(err)
		allUser = append(allUser, user)
	}
	rows.Close()
	return allUser
}

//GetUser database
func GetUser(employeeID string) User {
	db := getConnection()
	defer db.Close()
	sqlQuery := `
	SELECT APP_USER.ID,APP_USER.employee_id,APP_USER.username,APP_USER.password,APP_USER.name,APP_USER.sex,APP_USER.department,
	APP_USER.email, APP_USER.telephone,APP_USER.createDate,APP_USER.updateDate, APP_ROLE.role_name
	FROM APP_USER,APP_USERS_APP_ROLES,APP_ROLE 
	WHERE APP_USER.employee_id=? and APP_USERS_APP_ROLES.app_user = APP_USER.ID and APP_USERS_APP_ROLES.app_role = APP_ROLE.id
	`
	rowsAppUser, err := db.Query(sqlQuery, employeeID)
	checkErr(err)
	user := User{}
	for rowsAppUser.Next() {
		var roleName string
		err = rowsAppUser.Scan(&user.ID, &user.EmployeeID, &user.Username, &user.Password, &user.Name, &user.Sex, &user.Department, &user.Email, &user.Telephone, &user.CreateDate, &user.UpdateDate, &roleName)
		checkErr(err)
		user.Roles = append(user.Roles, roleName)
	}
	rowsAppUser.Close()
	return user
}

//GetUserByID database
func GetUserByID(userID int) User {
	db := getConnection()
	defer db.Close()
	sqlQuery := `
	SELECT APP_USER.ID,APP_USER.employee_id,APP_USER.username,APP_USER.password,APP_USER.name,APP_USER.sex,APP_USER.department,
	APP_USER.email, APP_USER.telephone,APP_USER.createDate,APP_USER.updateDate, APP_ROLE.role_name
	FROM APP_USER,APP_USERS_APP_ROLES,APP_ROLE 
	WHERE APP_USER.ID=? and APP_USERS_APP_ROLES.app_user = APP_USER.ID and APP_USERS_APP_ROLES.app_role = APP_ROLE.id
	`
	rowsAppUser, err := db.Query(sqlQuery, userID)
	checkErr(err)
	user := User{}
	for rowsAppUser.Next() {
		var roleName string
		err = rowsAppUser.Scan(&user.ID, &user.EmployeeID, &user.Username, &user.Password, &user.Name, &user.Sex, &user.Department, &user.Email, &user.Telephone, &user.CreateDate, &user.UpdateDate, &roleName)
		checkErr(err)
		user.Roles = append(user.Roles, roleName)
	}
	rowsAppUser.Close()
	return user
}

//GetByUsername database
func GetByUsername(username string) User {
	db := getConnection()
	defer db.Close()
	sqlQuery := `
	SELECT APP_USER.ID,APP_USER.employee_id,APP_USER.username,APP_USER.password,APP_USER.name,APP_USER.sex,APP_USER.department,
	APP_USER.email, APP_USER.telephone,APP_USER.createDate,APP_USER.updateDate, APP_ROLE.role_name
	FROM APP_USER,APP_USERS_APP_ROLES,APP_ROLE 
	WHERE APP_USER.username=? and APP_USERS_APP_ROLES.app_user = APP_USER.ID and APP_USERS_APP_ROLES.app_role = APP_ROLE.id
	`
	rowsAppUser, err := db.Query(sqlQuery, username)
	checkErr(err)
	user := User{}
	for rowsAppUser.Next() {
		var roleName string
		err = rowsAppUser.Scan(&user.ID, &user.EmployeeID, &user.Username, &user.Password, &user.Name, &user.Sex, &user.Department, &user.Email, &user.Telephone, &user.CreateDate, &user.UpdateDate, &roleName)
		checkErr(err)
		user.Roles = append(user.Roles, roleName)
	}
	rowsAppUser.Close()
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
