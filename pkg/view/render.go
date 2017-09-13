package view

import "net/http"

// Index render view
func Index(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("index.html"), w, data)
}

// Login render view
func Login(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("login.html"), w, data)
}

// FamsDashboard render view
func FamsDashboard(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("fams/dashboard.html"), w, data)
}

// FamsRequest render view
func FamsRequest(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("fams/request.html"), w, data)
}

// AdminCreateUser render view
func AdminCreateUser(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("admin/create.html"), w, data)
}

// AdminList render view
func AdminList(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("admin/list.html"), w, data)
}

// AdminUser render view
func AdminUser(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("admin/user.html"), w, data)
}

// AdminUpdateUser render view
func AdminUpdateUser(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("admin/update.html"), w, data)
}

// ImsDashboard render view
func ImsDashboard(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/dashboard.html"), w, data)
}

// ImsCreateMasterII render view
func ImsCreateMasterII(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/createMasterII.html"), w, data)
}

// ImsMasterIIDetail render view
func ImsMasterIIDetail(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/masteriiDetail.html"), w, data)
}

// ImsWaittingApproveMasterIIList render view
func ImsWaittingApproveMasterIIList(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/waittingApproveMasterIIList.html"), w, data)
}

// ImsApproveMasterIIList render view
func ImsApproveMasterIIList(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/approveMasterIIList.html"), w, data)
}

// ImsRejectMasterIIList render view
func ImsRejectMasterIIList(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/rejectMasterIIList.html"), w, data)
}

// ImsUpdateMasterII render view
func ImsUpdateMasterII(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/updateMasterII.html"), w, data)
}

// ImsWaittingApproveMasterII render view
func ImsWaittingApproveMasterII(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/waittingApproveMasterII.html"), w, data)
}

// ImsCreateUploadIIDataSearch render view
func ImsCreateUploadIIDataSearch(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/uploadIIDataSearch.html"), w, data)
}

// ImsUploadData render view
func ImsUploadData(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/uploadIIData.html"), w, data)
}

// ImsUploadDataDetail render view
func ImsUploadDataDetail(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/uploadIIDataDetail.html"), w, data)
}

// ImsCheckUploadData render view
func ImsCheckUploadData(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/checkUploadData.html"), w, data)
}

// ImsWaitLeaderCheckDataList render view
func ImsWaitLeaderCheckDataList(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/waittingLeaderCheckList.html"), w, data)
}

// ImsSearchReport render view
func ImsSearchReport(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/searchReport.html"), w, data)
}

// CustomerDashboard render view
func CustomerDashboard(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("customer/dashboard.html"), w, data)
}

// CustomerCreate render view
func CustomerCreate(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("customer/create.html"), w, data)
}
