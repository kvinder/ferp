package model

import (
	"fmt"
	"strings"
)

// InspectionData type
type InspectionData struct {
	ID                 int
	INNumber           string
	MasterII           MasterInspection
	Status             string
	WorkOrder          string
	Process            string
	TypeInspection     string
	QtyInspection      int
	CreateBy           User
	UpdateBy           User
	Remark             string
	CreateDate         string
	UpdateDate         string
	FileInspectionData FileUpload
	HistoryIN          []History
}

//CreateInspectionData database
func CreateInspectionData(inspectionData InspectionData) int {
	db := getConnection()
	defer db.Close()
	sqlQuery := `INSERT INTO InspectionData (INNumber,MasterII,FileInspectionData,workOrder,process,typeInspection,qtyInspection,
	Status,remark,CreateDate,UpdateDate,CreateBy,UpdateBy)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`
	result, err := db.Exec(sqlQuery,
		inspectionData.INNumber,
		inspectionData.MasterII.ID,
		inspectionData.FileInspectionData.ID,
		inspectionData.WorkOrder,
		inspectionData.Process,
		inspectionData.TypeInspection,
		inspectionData.QtyInspection,
		inspectionData.Status,
		inspectionData.Remark,
		inspectionData.CreateDate,
		inspectionData.UpdateDate,
		inspectionData.CreateBy.ID,
		inspectionData.UpdateBy.ID,
	)
	checkErr(err)
	lastID, _ := result.LastInsertId()
	intLastID := int(lastID)
	sqlQuery = `UPDATE InspectionData SET INNumber = ? WHERE id = ?`
	_, err = db.Exec(sqlQuery, fmt.Sprintf("IN%06d", intLastID), intLastID)
	checkErr(err)
	inspectionData.ID = intLastID
	history := History{
		Status:         inspectionData.Status,
		CreateBy:       inspectionData.CreateBy,
		Remark:         "upload inspection data",
		CreateDate:     inspectionData.CreateDate,
		InspecData:     inspectionData,
		FileInspecData: inspectionData.FileInspectionData,
	}
	CreateHistory(history)
	return inspectionData.ID
}

//UpdateInspectionData database
func UpdateInspectionData(inspectionData InspectionData) int {
	db := getConnection()
	defer db.Close()
	sqlQuery := `UPDATE InspectionData SET Status = ?, UpdateDate = ?, UpdateBy = ?
	WHERE id = ?`
	_, err := db.Exec(sqlQuery,
		inspectionData.Status,
		inspectionData.UpdateDate,
		inspectionData.UpdateBy.ID,
		inspectionData.ID,
	)
	checkErr(err)
	history := History{
		Status:         inspectionData.Status,
		CreateBy:       inspectionData.UpdateBy,
		Remark:         inspectionData.Remark,
		CreateDate:     inspectionData.UpdateDate,
		InspecData:     inspectionData,
		FileInspecData: inspectionData.FileInspectionData,
	}
	CreateHistory(history)
	return inspectionData.ID
}

//GetInspectionData database
func GetInspectionData(id int) InspectionData {
	db := getConnection()
	defer db.Close()
	sqlQuery := `SELECT ID,INNumber,MasterII,FileInspectionData,workOrder,process,typeInspection,
	qtyInspection,Status,remark,strftime('%Y-%m-%d %H:%M', CreateDate),UpdateDate,CreateBy,UpdateBy
	FROM InspectionData WHERE id = ?`
	rowsInspectionData, err := db.Query(sqlQuery, id)
	checkErr(err)
	inspecData := InspectionData{}
	var masterIIID, fileInspecID, createByID, updateByID int
	for rowsInspectionData.Next() {
		err = rowsInspectionData.Scan(
			&inspecData.ID,
			&inspecData.INNumber,
			&masterIIID,
			&fileInspecID,
			&inspecData.WorkOrder,
			&inspecData.Process,
			&inspecData.TypeInspection,
			&inspecData.QtyInspection,
			&inspecData.Status,
			&inspecData.Remark,
			&inspecData.CreateDate,
			&inspecData.UpdateDate,
			&createByID,
			&updateByID,
		)
		checkErr(err)
	}
	rowsInspectionData.Close()
	inspecData.MasterII = GetMasterII(masterIIID)
	if fileInspecID != 0 {
		inspecData.FileInspectionData = GetFileUploadByID(fileInspecID)
	}
	inspecData.CreateBy = GetUserByID(createByID)
	inspecData.UpdateBy = GetUserByID(updateByID)
	inspecData.HistoryIN = GetInspectionDataHistorys(inspecData.ID)
	return inspecData
}

//GetWorkOrderLike form database
func GetWorkOrderLike(workOrder string) []string {
	db := getConnection()
	defer db.Close()
	s := "%" + workOrder + "%"
	if len(workOrder) <= 0 {
		s = "%"
	}
	sqlQuery := "SELECT workOrder FROM InspectionData WHERE UPPER(workOrder) LIKE UPPER(?)"
	rows, err := db.Query(sqlQuery, s)
	checkErr(err)
	var allWorkOrder []string
	for rows.Next() {
		var wo string
		err = rows.Scan(&wo)
		checkErr(err)
		allWorkOrder = append(allWorkOrder, wo)
	}
	rows.Close()
	return allWorkOrder
}

//GetAllInspectionDataByStatus form database
func GetAllInspectionDataByStatus(status ...string) []InspectionData {
	db := getConnection()
	defer db.Close()
	stuff := make([]interface{}, len(status))
	for i, value := range status {
		stuff[i] = value
	}
	sqlQuery := `SELECT ID,INNumber,MasterII,FileInspectionData,workOrder,process,typeInspection,
	qtyInspection,Status,remark,strftime('%Y-%m-%d %H:%M', CreateDate),UpdateDate,CreateBy,UpdateBy
	FROM InspectionData WHERE status in (?` + strings.Repeat(",?", len(status)-1) + `)`
	rowsInspectionData, err := db.Query(sqlQuery, stuff...)
	checkErr(err)
	var inspectionDataList []InspectionData
	var masterIIID, fileInspecID, createByID, updateByID int
	for rowsInspectionData.Next() {
		inspecData := InspectionData{}
		err = rowsInspectionData.Scan(
			&inspecData.ID,
			&inspecData.INNumber,
			&masterIIID,
			&fileInspecID,
			&inspecData.WorkOrder,
			&inspecData.Process,
			&inspecData.TypeInspection,
			&inspecData.QtyInspection,
			&inspecData.Status,
			&inspecData.Remark,
			&inspecData.CreateDate,
			&inspecData.UpdateDate,
			&createByID,
			&updateByID,
		)
		checkErr(err)
		inspecData.MasterII = GetMasterII(masterIIID)
		if fileInspecID != 0 {
			inspecData.FileInspectionData = GetFileUploadByID(fileInspecID)
		}
		inspecData.CreateBy = GetUserByID(createByID)
		inspecData.UpdateBy = GetUserByID(updateByID)
		inspecData.HistoryIN = GetInspectionDataHistorys(inspecData.ID)
		inspectionDataList = append(inspectionDataList, inspecData)
	}
	rowsInspectionData.Close()
	return inspectionDataList
}

//GetAllInspectionDataBetweenDate form database
func GetAllInspectionDataBetweenDate(startDate, endDate string) []InspectionData {
	db := getConnection()
	defer db.Close()
	sqlQuery := `SELECT ID,INNumber,MasterII,FileInspectionData,workOrder,process,typeInspection,
	qtyInspection,Status,remark,strftime('%Y-%m-%d %H:%M', CreateDate),UpdateDate,CreateBy,UpdateBy
	FROM InspectionData WHERE CreateDate BETWEEN ? and ?`
	rowsInspectionData, err := db.Query(sqlQuery, startDate, endDate)
	checkErr(err)
	var inspectionDataList []InspectionData
	var masterIIID, fileInspecID, createByID, updateByID int
	for rowsInspectionData.Next() {
		inspecData := InspectionData{}
		err = rowsInspectionData.Scan(
			&inspecData.ID,
			&inspecData.INNumber,
			&masterIIID,
			&fileInspecID,
			&inspecData.WorkOrder,
			&inspecData.Process,
			&inspecData.TypeInspection,
			&inspecData.QtyInspection,
			&inspecData.Status,
			&inspecData.Remark,
			&inspecData.CreateDate,
			&inspecData.UpdateDate,
			&createByID,
			&updateByID,
		)
		checkErr(err)
		inspecData.MasterII = GetMasterII(masterIIID)
		if fileInspecID != 0 {
			inspecData.FileInspectionData = GetFileUploadByID(fileInspecID)
		}
		inspecData.CreateBy = GetUserByID(createByID)
		inspecData.UpdateBy = GetUserByID(updateByID)
		inspecData.HistoryIN = GetInspectionDataHistorys(inspecData.ID)
		inspectionDataList = append(inspectionDataList, inspecData)
	}
	rowsInspectionData.Close()
	return inspectionDataList
}

//GetAllInspectionData form database
func GetAllInspectionData() []InspectionData {
	db := getConnection()
	defer db.Close()
	sqlQuery := `SELECT ID,INNumber,MasterII,FileInspectionData,workOrder,process,typeInspection,
	qtyInspection,Status,remark, strftime('%Y-%m-%d %H:%M', CreateDate),UpdateDate,CreateBy,UpdateBy
	FROM InspectionData`
	rowsInspectionData, err := db.Query(sqlQuery)
	checkErr(err)
	var inspectionDataList []InspectionData
	var masterIIID, fileInspecID, createByID, updateByID int
	for rowsInspectionData.Next() {
		inspecData := InspectionData{}
		err = rowsInspectionData.Scan(
			&inspecData.ID,
			&inspecData.INNumber,
			&masterIIID,
			&fileInspecID,
			&inspecData.WorkOrder,
			&inspecData.Process,
			&inspecData.TypeInspection,
			&inspecData.QtyInspection,
			&inspecData.Status,
			&inspecData.Remark,
			&inspecData.CreateDate,
			&inspecData.UpdateDate,
			&createByID,
			&updateByID,
		)
		checkErr(err)
		inspecData.MasterII = GetMasterII(masterIIID)
		if fileInspecID != 0 {
			inspecData.FileInspectionData = GetFileUploadByID(fileInspecID)
		}
		inspecData.CreateBy = GetUserByID(createByID)
		inspecData.UpdateBy = GetUserByID(updateByID)
		inspecData.HistoryIN = GetInspectionDataHistorys(inspecData.ID)
		inspectionDataList = append(inspectionDataList, inspecData)
	}
	rowsInspectionData.Close()
	return inspectionDataList
}
