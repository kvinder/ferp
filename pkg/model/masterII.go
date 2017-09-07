package model

import (
	"fmt"
	"strings"
)

// MasterInspection type
type MasterInspection struct {
	ID              int
	MINumber        string
	CustomerMaterII Customer
	PartNumber      string
	PartName        string
	Revision        string
	Drawing         FileUpload
	Inspection      FileUpload
	File3           FileUpload
	File4           FileUpload
	File5           FileUpload
	TextFile3       string
	TextFile4       string
	TextFile5       string
	Status          string
	CreateDate      string
	UpdateDate      string
	Remark          string
	CreateBy        User
	UpdateBy        User
	HistoryMI       []History
}

//CreateMasterInspection create
func CreateMasterInspection(masterII MasterInspection) int {
	sqlQuery := `INSERT INTO MASTERINSPECTION (customer,partNumber,partName,revision,drawing,inspection,
	file3,file4,file5,textFile3,textFile4,textFile5,status,createDate,updateDate,createBy,updateBy)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	result, err := db.Exec(sqlQuery,
		masterII.CustomerMaterII.ID,
		masterII.PartNumber,
		masterII.PartName,
		masterII.Revision,
		masterII.Drawing.ID,
		masterII.Inspection.ID,
		masterII.File3.ID,
		masterII.File4.ID,
		masterII.File5.ID,
		masterII.TextFile3,
		masterII.TextFile4,
		masterII.TextFile5,
		masterII.Status,
		masterII.CreateDate,
		masterII.UpdateDate,
		masterII.CreateBy.ID,
		masterII.UpdateBy.ID,
	)
	checkErr(err)
	lastID, _ := result.LastInsertId()
	intLastID := int(lastID)
	sqlQuery = `UPDATE MASTERINSPECTION SET MINumber = ? WHERE id = ?`
	_, err = db.Exec(sqlQuery, fmt.Sprintf("MI%06d", intLastID), intLastID)
	checkErr(err)
	masterII.ID = intLastID
	history := History{
		MasterII:   masterII,
		Status:     masterII.Status,
		CreateBy:   masterII.CreateBy,
		Remark:     "create master inspection",
		CreateDate: masterII.CreateDate,
		Drawing:    masterII.Drawing,
		Inspection: masterII.Inspection,
		File3:      masterII.File3,
		File4:      masterII.File4,
		File5:      masterII.File5,
		TextFile3:  masterII.TextFile3,
		TextFile4:  masterII.TextFile4,
		TextFile5:  masterII.TextFile5,
	}
	CreateHistory(history)
	return masterII.ID
}

//UpdateMasterInspection create
func UpdateMasterInspection(masterII MasterInspection) int {
	sqlQuery := `UPDATE MASTERINSPECTION SET customer=?,partNumber=?,partName=?,revision=?,
	drawing=?,inspection=?,file3=?,file4=?,file5=?,textFile3=?,textFile4=?,textFile5=?,
	status=?,updateDate=?,updateBy=? WHERE id = ?`
	_, err := db.Exec(sqlQuery,
		masterII.CustomerMaterII.ID,
		masterII.PartNumber,
		masterII.PartName,
		masterII.Revision,
		masterII.Drawing.ID,
		masterII.Inspection.ID,
		masterII.File3.ID,
		masterII.File4.ID,
		masterII.File5.ID,
		masterII.TextFile3,
		masterII.TextFile4,
		masterII.TextFile5,
		masterII.Status,
		masterII.UpdateDate,
		masterII.UpdateBy.ID,
		masterII.ID,
	)
	checkErr(err)
	remarkHistory := ""
	if masterII.Status == "UPDATE_MASTER_II" {
		remarkHistory = "update master inspection"
	}
	if masterII.Status == "DELETE_MASTER_II" {
		remarkHistory = "delete master inspection"
	}
	if masterII.Status == "APPROVE_MASTER_II" {
		remarkHistory = "approve master inspection"
	}
	if masterII.Status == "REJECT_MASTER_II" {
		remarkHistory = masterII.Remark
	}
	history := History{
		MasterII:   masterII,
		Status:     masterII.Status,
		CreateBy:   masterII.UpdateBy,
		Remark:     remarkHistory,
		CreateDate: masterII.UpdateDate,
		Drawing:    masterII.Drawing,
		Inspection: masterII.Inspection,
		File3:      masterII.File3,
		File4:      masterII.File4,
		File5:      masterII.File5,
		TextFile3:  masterII.TextFile3,
		TextFile4:  masterII.TextFile4,
		TextFile5:  masterII.TextFile5,
	}
	CreateHistory(history)
	return masterII.ID
}

//GetMasterII getting
func GetMasterII(id int) MasterInspection {
	sqlQuery := `SELECT id,MINumber,customer,partNumber,partName,revision,drawing,inspection,
	file3,file4,file5,textFile3,textFile4,textFile5,status,createDate,updateDate,createBy,updateBy
	FROM MASTERINSPECTION WHERE id = ?`
	rowsMasterII, err := db.Query(sqlQuery, id)
	checkErr(err)
	masII := MasterInspection{}
	var customerID, drawingID, inspectionID, file3ID, file4ID, file5ID, crateByID, updateByID int
	for rowsMasterII.Next() {
		err = rowsMasterII.Scan(
			&masII.ID,
			&masII.MINumber,
			&customerID,
			&masII.PartNumber,
			&masII.PartName,
			&masII.Revision,
			&drawingID,
			&inspectionID,
			&file3ID,
			&file4ID,
			&file5ID,
			&masII.TextFile3,
			&masII.TextFile4,
			&masII.TextFile5,
			&masII.Status,
			&masII.CreateDate,
			&masII.UpdateDate,
			&crateByID,
			&updateByID,
		)
		checkErr(err)
	}
	rowsMasterII.Close()
	masII.CustomerMaterII = GetCustomerByID(customerID)
	if drawingID != 0 {
		masII.Drawing = GetFileUploadByID(drawingID)
	}
	if inspectionID != 0 {
		masII.Inspection = GetFileUploadByID(inspectionID)
	}
	if file3ID != 0 {
		masII.File3 = GetFileUploadByID(file3ID)
	}
	if file4ID != 0 {
		masII.File4 = GetFileUploadByID(file4ID)
	}
	if file5ID != 0 {
		masII.File5 = GetFileUploadByID(file5ID)
	}
	masII.CreateBy = GetUserByID(crateByID)
	masII.UpdateBy = GetUserByID(updateByID)
	masII.HistoryMI = GetHistorys(masII.ID)
	return masII
}

//GetAllMasterIIByStatusIn getting
func GetAllMasterIIByStatusIn(status ...string) []MasterInspection {
	stuff := make([]interface{}, len(status))
	for i, value := range status {
		stuff[i] = value
	}
	sqlQuery := `SELECT id,MINumber,customer,partNumber,partName,revision,drawing,inspection,
	file3,file4,file5,textFile3,textFile4,textFile5,status,createDate,updateDate,createBy,updateBy
	FROM MASTERINSPECTION WHERE status in (?` + strings.Repeat(",?", len(status)-1) + `)`
	rowsMasterII, err := db.Query(sqlQuery, stuff...)
	checkErr(err)
	var masterInspections []MasterInspection
	var customerID, drawingID, inspectionID, file3ID, file4ID, file5ID, crateByID, updateByID int
	for rowsMasterII.Next() {
		masII := MasterInspection{}
		err = rowsMasterII.Scan(
			&masII.ID,
			&masII.MINumber,
			&customerID,
			&masII.PartNumber,
			&masII.PartName,
			&masII.Revision,
			&drawingID,
			&inspectionID,
			&file3ID,
			&file4ID,
			&file5ID,
			&masII.TextFile3,
			&masII.TextFile4,
			&masII.TextFile5,
			&masII.Status,
			&masII.CreateDate,
			&masII.UpdateDate,
			&crateByID,
			&updateByID,
		)
		checkErr(err)

		masII.CustomerMaterII = GetCustomerByID(customerID)
		if drawingID != 0 {
			masII.Drawing = GetFileUploadByID(drawingID)
		}
		if inspectionID != 0 {
			masII.Inspection = GetFileUploadByID(inspectionID)
		}
		if file3ID != 0 {
			masII.File3 = GetFileUploadByID(file3ID)
		}
		if file4ID != 0 {
			masII.File4 = GetFileUploadByID(file4ID)
		}
		if file5ID != 0 {
			masII.File5 = GetFileUploadByID(file5ID)
		}
		masII.CreateBy = GetUserByID(crateByID)
		masII.UpdateBy = GetUserByID(updateByID)
		masterInspections = append(masterInspections, masII)
	}
	rowsMasterII.Close()

	return masterInspections
}
