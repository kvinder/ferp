package app

import (
	"encoding/json"
	"ferp/pkg/model"
	"ferp/pkg/view"
	"net/http"
	"time"
)

func customerDashboard(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"customer":           "active open",
		"customer_dashboard": "active",
		"customer_list":      model.GetAllCustomer(),
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	view.CustomerDashboard(w, data)
}

func customerCreate(w http.ResponseWriter, r *http.Request) {
	userOnLogin, err := model.UserOnLogin(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !checkRoles([]string{"Admin", "QA_LINE", "QA_FA", "QA_Engineer", "Sale_Co", "Sale_Out"}, userOnLogin.Roles) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		t := time.Now()
		now := t.Format("2006-01-02 15:04:05")
		customer := model.Customer{
			Name:       r.FormValue("inputCustomer"),
			TypeName:   r.FormValue("inputCustomerType"),
			CreateDate: now,
		}
		model.CreateCustomer(customer)
		http.Redirect(w, r, "/customer/dashboard", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"customer":        "active open",
		"customer_create": "active",
		"nameLogin":       userOnLogin.Name,
	}
	data = setAut(data, userOnLogin.Roles)
	view.CustomerCreate(w, data)
}

func findCustomersJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userOnLogin, err := model.UserOnLogin(r)
		if err != nil {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if !checkRoles([]string{"Admin", "QA_LINE", "QA_FA", "QA_Engineer", "Sale_Co", "Sale_Out"}, userOnLogin.Roles) {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		body := bodyToJSON(r)
		customers := model.GetCustomerLike(body["customer"])
		var customerArray []string
		for _, customer := range customers {
			customerArray = append(customerArray, customer.Name)
		}
		mapData := make(map[string][]string)
		mapData["customer"] = customerArray
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mapData)
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
