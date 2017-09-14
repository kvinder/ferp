package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var databaseName, databaseURL string

//InitConnection database
func InitConnection(dbName, dbURL string) {
	databaseName = dbName
	databaseURL = dbURL
	fmt.Println("Database connect...")
}

func getConnection() *sql.DB {
	dbConnect, err := sql.Open(databaseName, databaseURL)
	if err != nil {
		log.Fatalf("can not connect database : %v", err)
	}
	return dbConnect
}

//CreateDatabase sql
func CreateDatabase() {
	sqlCreateTable := `
		CREATE TABLE APP_USER (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		employee_id varchar(255) NOT NULL UNIQUE,
		username varchar(255) DEFAULT NULL,
		password varchar(255) DEFAULT NULL,
		name varchar(255) DEFAULT NULL,
		sex varchar(255) DEFAULT NULL,
		department varchar(255) DEFAULT NULL,
		email varchar(255) DEFAULT NULL,
		telephone varchar(255) DEFAULT NULL,
		createDate datetime DEFAULT NULL,
		updateDate datetime DEFAULT NULL);
	  
		CREATE TABLE APP_ROLE (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		role_name varchar(255) DEFAULT NULL,
		role_description varchar(255) DEFAULT NULL);

		INSERT INTO APP_ROLE (role_name,role_description) VALUES('Sale_Co','Sale Co');
		INSERT INTO APP_ROLE (role_name,role_description) VALUES('Sale_Out','Sale Out');
		INSERT INTO APP_ROLE (role_name,role_description) VALUES('QA_FA','QA FA');
		INSERT INTO APP_ROLE (role_name,role_description) VALUES('QA_Engineer','QA Engineer');
		INSERT INTO APP_ROLE (role_name,role_description) VALUES('Purchase','Purchase');
		INSERT INTO APP_ROLE (role_name,role_description) VALUES('Engineer','Engineer');
		INSERT INTO APP_ROLE (role_name,role_description) VALUES('Admin','Admin');
		INSERT INTO APP_ROLE (role_name,role_description) VALUES('QA_LINE','QA Line');
		INSERT INTO APP_ROLE (role_name,role_description) VALUES('QA_OFFICE','QA Office');

		CREATE TABLE APP_USERS_APP_ROLES (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		app_user INTEGER,
		app_role INTEGER);

		INSERT INTO APP_USER (employee_id,username,password,name,sex,department,email,telephone,createDate,updateDate)
		VALUES('00000','admin',X'24326124313024796733756E5673676653744D6D71616A4E6D30484875566B486273435A58354B43626570676B7A65416673435A4B4F623538335932','Admin','Male','MIS','a.c@g.v','0000000000','2017-08-18 12:41:32','2017-08-18 14:06:31');

		INSERT INTO APP_USERS_APP_ROLES (app_user,app_role) VALUES(1,7);
		
		CREATE TABLE CUSTOMER (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		name varchar(255) DEFAULT NULL,
		typeName varchar(255) DEFAULT NULL,
		createDate datetime DEFAULT NULL);

		CREATE TABLE FileUpload (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		filename varchar(255) DEFAULT NULL,
		ContentType varchar(255) DEFAULT NULL,
		createDate datetime DEFAULT NULL);

		CREATE TABLE MASTERINSPECTION (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		MINumber varchar(255) DEFAULT NULL,
		Customer INTEGER DEFAULT NULL,
		partNumber varchar(255) DEFAULT NULL,
		partName varchar(255) DEFAULT NULL,
		revision varchar(255) DEFAULT NULL,
		drawing INTEGER DEFAULT NULL,
		Inspection INTEGER DEFAULT NULL,
		file3 INTEGER DEFAULT NULL,
		file4 INTEGER DEFAULT NULL,
		file5 INTEGER DEFAULT NULL,
		textFile3 varchar(255) DEFAULT NULL,
		textFile4 varchar(255) DEFAULT NULL,
		textFile5 varchar(255) DEFAULT NULL,
		Status varchar(255) DEFAULT NULL,
		CreateDate datetime DEFAULT NULL,
		UpdateDate datetime DEFAULT NULL,
		CreateBy INTEGER DEFAULT NULL,
		UpdateBy INTEGER DEFAULT NULL);

		CREATE TABLE InspectionData (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		INNumber varchar(255) DEFAULT NULL,
		MasterII INTEGER DEFAULT NULL,
		FileInspectionData INTEGER DEFAULT NULL,
		workOrder varchar(255) DEFAULT NULL,
		process varchar(255) DEFAULT NULL,
		typeInspection varchar(255) DEFAULT NULL,
		qtyInspection INTEGER DEFAULT NULL,
		Status varchar(255) DEFAULT NULL,
		remark varchar(255) DEFAULT NULL,
		CreateDate datetime DEFAULT NULL,
		UpdateDate datetime DEFAULT NULL,
		CreateBy INTEGER DEFAULT NULL,
		UpdateBy INTEGER DEFAULT NULL);

		CREATE TABLE History (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		MasterII INTEGER DEFAULT NULL,
		InspectionData INTEGER DEFAULT NULL,
		status varchar(255) DEFAULT NULL,
		CreateBy INTEGER DEFAULT NULL,
		Remark varchar(255) DEFAULT NULL,
		createDate datetime DEFAULT NULL,
		Drawing INTEGER DEFAULT NULL,
		Inspection INTEGER DEFAULT NULL,
		file3 INTEGER DEFAULT NULL,
		file4 INTEGER DEFAULT NULL,
		file5 INTEGER DEFAULT NULL,
		textFile3 varchar(255) DEFAULT NULL,
		textFile4 varchar(255) DEFAULT NULL,
		textFile5 varchar(255) DEFAULT NULL,
		FileInspectionData INTEGER DEFAULT NULL);
	  `
	db := getConnection()
	defer db.Close()
	_, err := db.Exec(sqlCreateTable)
	checkErr(err)
	fmt.Println("Create Table...")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
