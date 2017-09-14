package view

import "net/http"

// Index render view
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}

// Login render view
func Login(w http.ResponseWriter, data interface{}) {
	render(tpLogin, w, data)
}

// FamsDashboard render view
func FamsDashboard(w http.ResponseWriter, data interface{}) {
	render(tpFamsDashboard, w, data)
}

// FamsRequest render view
func FamsRequest(w http.ResponseWriter, data interface{}) {
	render(tpFamsRequest, w, data)
}

// AdminCreateUser render view
func AdminCreateUser(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreateUser, w, data)
}

// AdminList render view
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminList, w, data)
}

// AdminUser render view
func AdminUser(w http.ResponseWriter, data interface{}) {
	render(tpAdminUser, w, data)
}

// AdminUpdateUser render view
func AdminUpdateUser(w http.ResponseWriter, data interface{}) {
	render(tpAdminUpdateUser, w, data)
}

// ImsDashboard render view
func ImsDashboard(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("ims/dashboard.html"), w, data)
}

// ImsCreateMasterII render view
func ImsCreateMasterII(w http.ResponseWriter, data interface{}) {
	render(tpImsCreateMasterII, w, data)
}

// ImsMasterIIDetail render view
func ImsMasterIIDetail(w http.ResponseWriter, data interface{}) {
	render(tpImsMasterIIDetail, w, data)
}

// ImsWaittingApproveMasterIIList render view
func ImsWaittingApproveMasterIIList(w http.ResponseWriter, data interface{}) {
	render(tpImsWaittingApproveMasterIIList, w, data)
}

// ImsApproveMasterIIList render view
func ImsApproveMasterIIList(w http.ResponseWriter, data interface{}) {
	render(tpImsApproveMasterIIList, w, data)
}

// ImsRejectMasterIIList render view
func ImsRejectMasterIIList(w http.ResponseWriter, data interface{}) {
	render(tpImsRejectMasterIIList, w, data)
}

// ImsUpdateMasterII render view
func ImsUpdateMasterII(w http.ResponseWriter, data interface{}) {
	render(tpImsUpdateMasterII, w, data)
}

// ImsWaittingApproveMasterII render view
func ImsWaittingApproveMasterII(w http.ResponseWriter, data interface{}) {
	render(tpImsWaittingApproveMasterII, w, data)
}

// ImsCreateUploadIIDataSearch render view
func ImsCreateUploadIIDataSearch(w http.ResponseWriter, data interface{}) {
	render(tpImsCreateUploadIIDataSearch, w, data)
}

// ImsUploadData render view
func ImsUploadData(w http.ResponseWriter, data interface{}) {
	render(tpImsUploadData, w, data)
}

// ImsUploadDataDetail render view
func ImsUploadDataDetail(w http.ResponseWriter, data interface{}) {
	render(tpImsUploadDataDetail, w, data)
}

// ImsCheckUploadData render view
func ImsCheckUploadData(w http.ResponseWriter, data interface{}) {
	render(tpImsCheckUploadData, w, data)
}

// ImsWaitLeaderCheckDataList render view
func ImsWaitLeaderCheckDataList(w http.ResponseWriter, data interface{}) {
	render(tpImsWaitLeaderCheckDataList, w, data)
}

// ImsSearchReport render view
func ImsSearchReport(w http.ResponseWriter, data interface{}) {
	render(tpImsSearchReport, w, data)
}

// CustomerDashboard render view
func CustomerDashboard(w http.ResponseWriter, data interface{}) {
	render(tpCustomerDashboard, w, data)
}

// CustomerCreate render view
func CustomerCreate(w http.ResponseWriter, data interface{}) {
	render(tpCustomerCreate, w, data)
}
