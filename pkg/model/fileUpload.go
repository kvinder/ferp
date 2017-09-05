package model

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

//FileUpload type
type FileUpload struct {
	ID          int
	FileName    string
	ContentType string
	CreateDate  string
}

//UploadFile database
func UploadFile(r *http.Request, fileForm string) (fileupload FileUpload, err error) {
	drawingFile, typeHeader, err := r.FormFile(fileForm)
	file := FileUpload{}
	if drawingFile == nil || typeHeader == nil {
		return file, nil
	}
	if err != nil {
		return file, err
	}

	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	sqlInsertFile := `INSERT INTO FileUpload (filename,ContentType,createDate) VALUES (?,?,?)`
	res, err := db.Exec(sqlInsertFile, typeHeader.Filename, typeHeader.Header.Get("Content-Type"), now)
	checkErr(err)
	i, _ := res.LastInsertId()
	file.ID = int(i)
	file.FileName = typeHeader.Filename
	file.ContentType = typeHeader.Header.Get("Content-Type")
	mkdirPath := "./fileUploads/" + strconv.Itoa(file.ID)
	os.MkdirAll(mkdirPath, os.ModePerm)
	drFile, err := os.OpenFile("./fileUploads/"+strconv.Itoa(file.ID)+"/"+typeHeader.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	checkErr(err)
	_, err = io.Copy(drFile, drawingFile)
	if err != nil {
		sqlDelete := `DELETE FROM FileUpload WHERE ID = ?`
		_, err = db.Exec(sqlDelete, file.ID)
		checkErr(err)
		return file, err
	}
	defer drFile.Close()
	defer drawingFile.Close()
	return file, nil
}

//GetFileUploadByID getting
func GetFileUploadByID(fileUploadID int) FileUpload {
	sqlQuery := `SELECT id,filename,ContentType,createDate
	FROM FileUpload WHERE id = ?`
	rowsFileUpload, err := db.Query(sqlQuery, fileUploadID)
	checkErr(err)
	fileUpload := FileUpload{}
	for rowsFileUpload.Next() {
		err = rowsFileUpload.Scan(
			&fileUpload.ID,
			&fileUpload.FileName,
			&fileUpload.ContentType,
			&fileUpload.CreateDate,
		)
	}
	rowsFileUpload.Close()
	return fileUpload
}
