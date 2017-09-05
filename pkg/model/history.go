package model

// History type
type History struct {
	ID         int
	MasterII   MasterInspection
	Status     string
	CreateBy   User
	Remark     string
	CreateDate string
	Drawing    FileUpload
	Inspection FileUpload
	File3      FileUpload
	File4      FileUpload
	File5      FileUpload
	TextFile3  string
	TextFile4  string
	TextFile5  string
}

//CreateHistory database
func CreateHistory(history History) {
	sqlQuery := `INSERT INTO History (MasterII,status,createBy,Remark,
	createDate,Drawing,Inspection,file3,file4,file5,textFile3,textFile4,textFile5)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sqlQuery,
		history.MasterII.ID,
		history.Status,
		history.CreateBy.ID,
		history.Remark,
		history.CreateDate,
		history.Drawing.ID,
		history.Inspection.ID,
		history.File3.ID,
		history.File4.ID,
		history.File5.ID,
		history.TextFile3,
		history.TextFile4,
		history.TextFile5,
	)
	checkErr(err)
}

//GetHistorys database
func GetHistorys(masteriiID int) {
	sqlQuery := `SELECT id,status,CreateBy,Remark,createDate,Drawing,Inspection,
	file3,file4,file5,textFile3,textFile4,textFile5
	FROM History WHERE MasterII = ?`
	rowsHistorys, err := db.Query(sqlQuery, masteriiID)
	checkErr(err)
	var historys []History
	for rowsHistorys.Next() {
		h := History{}
		var createByID, drawingID, inspectionID, file3ID, file4ID, file5ID int
		err = rowsHistorys.Scan(
			&h.ID,
			&h.Status,
			&createByID,
			&h.Remark,
			&h.CreateDate,
			&drawingID,
			&inspectionID,
			&file3ID,
			&file4ID,
			&file5ID,
			&h.TextFile3,
			&h.TextFile4,
			&h.TextFile5,
		)
		checkErr(err)
		h.CreateBy = GetUserByID(createByID)
		if drawingID != 0 {
			h.Drawing = GetFileUploadByID(drawingID)
		}
		if inspectionID != 0 {
			h.Inspection = GetFileUploadByID(inspectionID)
		}
		if file3ID != 0 {
			h.File3 = GetFileUploadByID(file3ID)
		}
		if file4ID != 0 {
			h.File4 = GetFileUploadByID(file4ID)
		}
		if file5ID != 0 {
			h.File5 = GetFileUploadByID(file5ID)
		}
		historys = append(historys, h)
	}
	rowsHistorys.Close()
}
