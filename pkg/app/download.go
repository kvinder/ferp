package app

import (
	"ferp/pkg/model"
	"io"
	"net/http"
	"os"
	"strconv"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Query().Get("file")
	id, _ := strconv.Atoi(fileID)
	fileUpload := model.GetFileUploadByID(id)
	out, err := os.Open("./fileUploads/" + fileID + "/" + fileUpload.FileName)
	checkErr(err)
	defer out.Close()
	w.Header().Set("Content-Disposition", "inline; filename="+fileUpload.FileName)
	w.Header().Set("Content-Type", fileUpload.ContentType)
	io.Copy(w, out)
}
