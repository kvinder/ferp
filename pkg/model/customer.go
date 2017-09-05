package model

// Customer type
type Customer struct {
	ID         int
	Name       string
	TypeName   string
	CreateDate string
	Parts      []Part
}

//CreateCustomer database
func CreateCustomer(customer Customer) {
	sqlInsertCustomer := `INSERT INTO CUSTOMER (name,typeName,createDate) VALUES (UPPER(?),?,?)`
	_, err := db.Exec(sqlInsertCustomer, customer.Name, customer.TypeName, customer.CreateDate)
	checkErr(err)
}

//GetAllCustomer form database
func GetAllCustomer() []Customer {
	rows, err := db.Query("SELECT id, name, typeName, createDate FROM CUSTOMER")
	checkErr(err)
	var allCustomer []Customer
	for rows.Next() {
		customer := Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.TypeName, &customer.CreateDate)
		checkErr(err)
		allCustomer = append(allCustomer, customer)
	}
	rows.Close()
	return allCustomer
}

//GetCustomerByName form database
func GetCustomerByName(customer string) Customer {
	sqlQuery := "SELECT id, name, typeName, createDate FROM CUSTOMER WHERE UPPER(name) = UPPER(?)"
	rows, err := db.Query(sqlQuery, customer)
	checkErr(err)
	var customerQuery Customer
	for rows.Next() {
		err = rows.Scan(&customerQuery.ID, &customerQuery.Name, &customerQuery.TypeName, &customerQuery.CreateDate)
		checkErr(err)
	}
	rows.Close()
	return customerQuery
}

//GetCustomerByID form database
func GetCustomerByID(customerID int) Customer {
	sqlQuery := "SELECT id, name, typeName, createDate FROM CUSTOMER WHERE id = ?"
	rows, err := db.Query(sqlQuery, customerID)
	checkErr(err)
	var customerQuery Customer
	for rows.Next() {
		err = rows.Scan(&customerQuery.ID, &customerQuery.Name, &customerQuery.TypeName, &customerQuery.CreateDate)
		checkErr(err)
	}
	rows.Close()
	return customerQuery
}

//GetCustomerLike form database
func GetCustomerLike(customer string) []Customer {
	s := "%" + customer + "%"
	if len(customer) <= 0 {
		s = "%"
	}
	sqlQuery := "SELECT id, name, typeName, createDate FROM CUSTOMER WHERE UPPER(name) LIKE UPPER(?)"
	rows, err := db.Query(sqlQuery, s)
	checkErr(err)
	var allCustomer []Customer
	for rows.Next() {
		customer := Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.TypeName, &customer.CreateDate)
		checkErr(err)
		allCustomer = append(allCustomer, customer)
	}
	rows.Close()
	return allCustomer
}
