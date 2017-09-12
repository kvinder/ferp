package model

import "fmt"

// InspectionData type
type InspectionData struct {
	ID                 int
	INNumber           string
	MasterII           MasterInspection
	Status             string
	WorkOrder          string
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
	sqlQuery := `INSERT INTO InspectionData (INNumber,MasterII,FileInspectionData,workOrder,typeInspection,qtyInspection,
	Status,remark,CreateDate,UpdateDate,CreateBy,UpdateBy)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
	result, err := db.Exec(sqlQuery,
		inspectionData.INNumber,
		inspectionData.MasterII.ID,
		inspectionData.FileInspectionData.ID,
		inspectionData.WorkOrder,
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

//GetInspectionData database
func GetInspectionData(id int) InspectionData {
	sqlQuery := `SELECT ID,INNumber,MasterII,FileInspectionData,workOrder,typeInspection,
	qtyInspection,Status,remark,CreateDate,UpdateDate,CreateBy,UpdateBy
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
