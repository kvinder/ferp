package app

import (
	"encoding/json"
	"ferp/pkg/model"
	"ferp/pkg/view"
	"net/http"
	"strconv"
	"time"
)

func imsDashboard(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims":           "active open",
		"ims_dashboard": "active",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	t := time.Now()
	year, month, day := t.Date()
	startDateTime := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	endDateTime := time.Date(year, month, day, 23, 59, 59, 0, t.Location())
	startDateStr := startDateTime.Format("2006-01-02 15:04:05")
	endDateStr := endDateTime.Format("2006-01-02 15:04:05")
	data["inspectionDataList"] = model.GetAllInspectionDataBetweenDate(startDateStr, endDateStr)
	view.ImsDashboard(w, data)
}

func imsCreateMasterII(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
		"ims_create_master_ii": "active",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}

	checkAuthorization([]string{"Admin", "QA_OFFICE", "QA_Engineer", "QA_FA"}, userOnLogin, w, r)
	if r.Method == http.MethodPost {
		fileDrawing, err := model.UploadFile(r, "id-input-file-1")
		checkErr(err)
		inspection, err := model.UploadFile(r, "id-input-file-2")
		checkErr(err)
		file3, err := model.UploadFile(r, "id-input-file-3")
		checkErr(err)
		file4, err := model.UploadFile(r, "id-input-file-4")
		checkErr(err)
		file5, err := model.UploadFile(r, "id-input-file-5")
		checkErr(err)
		customer := model.GetCustomerByName(r.FormValue("inputCustomer"))
		t := time.Now()
		now := t.Format("2006-01-02 15:04:05")
		masterII := model.MasterInspection{
			CustomerMaterII: customer,
			PartNumber:      r.FormValue("inputPartNo"),
			PartName:        r.FormValue("inputPartName"),
			Revision:        r.FormValue("inputRevision"),
			Drawing:         fileDrawing,
			Inspection:      inspection,
			File3:           file3,
			File4:           file4,
			File5:           file5,
			TextFile3:       r.FormValue("inputFile1"),
			TextFile4:       r.FormValue("inputFile2"),
			TextFile5:       r.FormValue("inputFile3"),
			Status:          "CREATE_MASTER_II",
			CreateDate:      now,
			UpdateDate:      now,
			CreateBy:        userOnLogin,
			UpdateBy:        userOnLogin,
		}
		id := model.CreateMasterInspection(masterII)
		http.Redirect(w, r, "/ims/masterii?detail="+strconv.Itoa(id), http.StatusSeeOther)
		return
	}
	view.ImsCreateMasterII(w, data)
}

func imsUpdateMasterII(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	checkAuthorization([]string{"Admin", "QA_OFFICE", "QA_Engineer", "QA_FA"}, userOnLogin, w, r)
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		i, _ := strconv.Atoi(id)
		currentMasterII := model.GetMasterII(i)
		currentFile1 := r.FormValue("currentFile1")
		if currentFile1 != "yes" {
			currentMasterII.Drawing, err = model.UploadFile(r, "id-input-file-1")
			checkErr(err)
		}
		currentFile2 := r.FormValue("currentFile2")
		if currentFile2 != "yes" {
			currentMasterII.Inspection, err = model.UploadFile(r, "id-input-file-2")
			checkErr(err)
		}
		currentFile3 := r.FormValue("currentFile3")
		if currentFile3 != "yes" {
			currentMasterII.File3, err = model.UploadFile(r, "id-input-file-3")
			checkErr(err)
		}
		currentFile4 := r.FormValue("currentFile4")
		if currentFile4 != "yes" {
			currentMasterII.File4, err = model.UploadFile(r, "id-input-file-4")
			checkErr(err)
		}
		currentFile5 := r.FormValue("currentFile5")
		if currentFile5 != "yes" {
			currentMasterII.File5, err = model.UploadFile(r, "id-input-file-5")
			checkErr(err)
		}
		customer := model.GetCustomerByName(r.FormValue("inputCustomer"))
		t := time.Now()
		now := t.Format("2006-01-02 15:04:05")
		currentMasterII.CustomerMaterII = customer
		currentMasterII.PartNumber = r.FormValue("inputPartName")
		currentMasterII.PartName = r.FormValue("inputPartName")
		currentMasterII.Revision = r.FormValue("inputRevision")
		currentMasterII.TextFile3 = r.FormValue("inputFile1")
		currentMasterII.TextFile4 = r.FormValue("inputFile2")
		currentMasterII.TextFile5 = r.FormValue("inputFile3")
		currentMasterII.Status = r.FormValue("statusInput")
		currentMasterII.UpdateDate = now
		currentMasterII.UpdateBy = userOnLogin
		idRes := model.UpdateMasterInspection(currentMasterII)
		http.Redirect(w, r, "/ims/masterii?detail="+strconv.Itoa(idRes), http.StatusSeeOther)
		return
	}
	view.ImsCreateMasterII(w, data)
}

func imsMasterII(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
	}
	detailID := r.URL.Query().Get("detail")
	updateID := r.URL.Query().Get("update")
	waitingapproveID := r.URL.Query().Get("waitingapprove")
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
		if len(detailID) == 0 && len(waitingapproveID) == 0 && len(updateID) != 0 {
			checkAuthorization([]string{"Admin", "QA_OFFICE", "QA_Engineer", "QA_FA"}, userOnLogin, w, r)
			i, _ := strconv.Atoi(updateID)
			data["detail"] = model.GetMasterII(i)
			view.ImsUpdateMasterII(w, data)
			return
		}
		if len(detailID) == 0 && len(waitingapproveID) != 0 && len(updateID) == 0 {
			checkAuthorization([]string{"Admin", "QA_Engineer", "QA_FA"}, userOnLogin, w, r)
			i, _ := strconv.Atoi(waitingapproveID)
			data["detail"] = model.GetMasterII(i)
			view.ImsWaittingApproveMasterII(w, data)
			return
		}
	}
	if len(detailID) != 0 && len(updateID) == 0 {
		i, _ := strconv.Atoi(detailID)
		data["detail"] = model.GetMasterII(i)
		view.ImsMasterIIDetail(w, data)
		return
	}
	http.NotFound(w, r)
	return
}

func imsMasterIIApproveOrReject(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	checkAuthorization([]string{"Admin", "QA_Engineer", "QA_FA"}, userOnLogin, w, r)
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		i, _ := strconv.Atoi(id)
		currentMasterII := model.GetMasterII(i)
		currentMasterII.Status = r.FormValue("approveOrReject")
		currentMasterII.Remark = r.FormValue("inputReasonReject")
		currentMasterII.UpdateBy = userOnLogin
		t := time.Now()
		now := t.Format("2006-01-02 15:04:05")
		currentMasterII.UpdateDate = now
		idRes := model.UpdateMasterInspection(currentMasterII)
		http.Redirect(w, r, "/ims/masterii?detail="+strconv.Itoa(idRes), http.StatusSeeOther)
		return
	}
}

func imsWaittingApproveMasterIIList(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
		"ims_waitting_approve_master_ii": "active",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	data["masterIIListWaitApprove"] = model.GetAllMasterIIByStatusIn("CREATE_MASTER_II", "UPDATE_MASTER_II")
	view.ImsWaittingApproveMasterIIList(w, data)
}

func imsApproveMasterIIList(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
		"ims_approve_master_ii": "active",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	data["masterIIApproveList"] = model.GetAllMasterIIByStatusIn("APPROVE_MASTER_II")
	view.ImsApproveMasterIIList(w, data)
}

func imsRejectMasterIIList(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
		"ims_reject_master_ii": "active",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	data["masterIIRejectList"] = model.GetAllMasterIIByStatusIn("REJECT_MASTER_II")
	view.ImsRejectMasterIIList(w, data)
}

func imsUploadIIDataSearch(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims":                "active open",
		"ims_upload_ii_data": "active",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	checkAuthorization([]string{"Admin", "QA_OFFICE", "QA_Engineer", "QA_FA", "QA_LINE"}, userOnLogin, w, r)
	data["masterIIApproveList"] = model.GetAllMasterIIByStatusIn("APPROVE_MASTER_II")
	view.ImsCreateUploadIIDataSearch(w, data)
}

func imsUploadData(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	checkAuthorization([]string{"Admin", "QA_OFFICE", "QA_Engineer", "QA_FA", "QA_LINE"}, userOnLogin, w, r)
	if r.Method == http.MethodPost {
		fileData, err := model.UploadFile(r, "id-input-file")
		checkErr(err)
		t := time.Now()
		now := t.Format("2006-01-02 15:04:05")
		idMasterStr := r.FormValue("idMasterII")
		idMasterII, _ := strconv.Atoi(idMasterStr)
		qtyInspec, _ := strconv.Atoi(r.FormValue("inputQty"))
		inspectionData := model.InspectionData{
			CreateDate:         now,
			UpdateDate:         now,
			CreateBy:           userOnLogin,
			UpdateBy:           userOnLogin,
			FileInspectionData: fileData,
			MasterII:           model.GetMasterII(idMasterII),
			Status:             "UPLOAD_INSPECTION_DATA",
			WorkOrder:          r.FormValue("inputWorkOrder"),
			TypeInspection:     r.FormValue("field-type"),
			QtyInspection:      qtyInspec,
			Remark:             r.FormValue("id-remark"),
		}
		idRes := model.CreateInspectionData(inspectionData)
		http.Redirect(w, r, "/ims/datainspection?detail="+strconv.Itoa(idRes), http.StatusSeeOther)
		return
	}
	masterid := r.URL.Query().Get("masterid")
	i, _ := strconv.Atoi(masterid)
	data["detail"] = model.GetMasterII(i)
	view.ImsUploadData(w, data)
}

func imsUploadDataDetail(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	detailid := r.URL.Query().Get("detail")
	i, _ := strconv.Atoi(detailid)
	data["detail"] = model.GetInspectionData(i)
	view.ImsUploadDataDetail(w, data)
}

func imsCheckUploadData(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	checkAuthorization([]string{"Admin", "QA_Engineer", "QA_FA", "QA_OFFICE"}, userOnLogin, w, r)
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		i, _ := strconv.Atoi(id)
		inspectionData := model.GetInspectionData(i)
		t := time.Now()
		now := t.Format("2006-01-02 15:04:05")
		inspectionData.Remark = r.FormValue("inputConfirm")
		inspectionData.UpdateBy = userOnLogin
		inspectionData.UpdateDate = now
		if inspectionData.Status != "UPLOAD_INSPECTION_DATA" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		inspectionData.Status = r.FormValue("checked-input")
		model.UpdateInspectionData(inspectionData)
		http.Redirect(w, r, "/ims/waitleaderchecklist", http.StatusSeeOther)
		return
	}
	detailid := r.URL.Query().Get("check")
	i, _ := strconv.Atoi(detailid)
	data["detail"] = model.GetInspectionData(i)
	view.ImsCheckUploadData(w, data)
}

func imsSearchReportList(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims":               "active open",
		"ims_search_report": "active",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	data["inspectionDataList"] = model.GetAllInspectionData()
	view.ImsSearchReport(w, data)
}

func findWorkOrdersJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userOnLogin, err := model.UserOnLogin(r)
		if err != nil {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		checkAuthorization([]string{"Admin", "QA_OFFICE", "QA_Engineer", "QA_FA", "QA_LINE"}, userOnLogin, w, r)
		body := bodyToJSON(r)
		workOrders := model.GetWorkOrderLike(body["workOrder"])
		var workOrderArray []string
		for _, workOrder := range workOrders {
			workOrderArray = append(workOrderArray, workOrder)
		}
		mapData := make(map[string][]string)
		mapData["workOrders"] = workOrderArray
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mapData)
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}

func imsWaittingLeaderCheckDataList(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ims": "active open",
		"ims_waitting_leader_check_data": "active",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	data["inspectionDataList"] = model.GetAllInspectionDataByStatus("UPLOAD_INSPECTION_DATA")
	view.ImsWaitLeaderCheckDataList(w, data)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
