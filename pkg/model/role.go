package model

// Role type
type Role struct {
	ID              int
	RoleName        string
	RoleDescription string
}

//ListRoles database
func ListRoles() []Role {
	rows, err := db.Query("SELECT ID, role_name, role_description FROM APP_ROLE")
	checkErr(err)
	var allRole []Role
	for rows.Next() {
		role := Role{}
		err = rows.Scan(&role.ID, &role.RoleName, &role.RoleDescription)
		checkErr(err)
		allRole = append(allRole, role)
	}
	rows.Close()
	return allRole
}
