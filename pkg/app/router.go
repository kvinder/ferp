package app

import (
	"encoding/json"
	"ferp/pkg/model"
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
)

// Router handlers to mux
func Router(mux *http.ServeMux) {
	mux.Handle("/-/", http.StripPrefix("/-", http.FileServer(http.Dir("static"))))

	midelwareMux := http.NewServeMux()
	midelwareMux.HandleFunc("/", index)
	midelwareMux.HandleFunc("/login", login)
	midelwareMux.HandleFunc("/logout", logout)
	midelwareMux.HandleFunc("/fa/dashboard", famsDashboard)
	midelwareMux.HandleFunc("/ims/dashboard", imsDashboard)
	midelwareMux.HandleFunc("/ims/waittingapproveii", imsWaittingApproveMasterIIList)
	midelwareMux.HandleFunc("/ims/approveiilist", imsApproveMasterIIList)
	midelwareMux.HandleFunc("/ims/rejectiilist", imsRejectMasterIIList)

	//Admin
	midelwareMux.HandleFunc("/admin/create", adminCreate)
	midelwareMux.HandleFunc("/admin/list", adminList)
	midelwareMux.HandleFunc("/admin/user", adminUser)
	midelwareMux.HandleFunc("/admin/user/update", adminUpdate)

	//FAMS
	midelwareMux.HandleFunc("/fams/request", famsRequest)

	//ims
	midelwareMux.HandleFunc("/ims/createmasterii", imsCreateMasterII)
	midelwareMux.HandleFunc("/ims/masterii", imsMasterII)
	midelwareMux.HandleFunc("/ims/updatemasterii", imsUpdateMasterII)
	midelwareMux.HandleFunc("/ims/approverejectmasterii", imsMasterIIApproveOrReject)

	//fileUpload
	midelwareMux.HandleFunc("/file", downloadFile)

	//customer
	midelwareMux.HandleFunc("/customer/dashboard", customerDashboard)
	midelwareMux.HandleFunc("/customer/create", customerCreate)
	//customer JSON
	midelwareMux.HandleFunc("/customer/findcustomer", findCustomersJSON)

	mux.Handle("/", midelware(midelwareMux))

}

var pathNoNeedLogin = map[string]bool{
	"/":                          true,
	"/login":                     true,
	"/fa/dashboard":              true,
	"/file":                      true,
	"/ims/masterii":              true,
	"/ims/waittingapproveii":     true,
	"/ims/approveiilist":         true,
	"/ims/rejectiilist":          true,
	"/admin/create":              false,
	"/admin/list":                false,
	"/admin/user":                false,
	"/admin/user/update":         false,
	"/fams/request":              false,
	"/ims/dashboard":             true,
	"/ims/createmasterii":        false,
	"/ims/updatemasterii":        false,
	"/ims/approverejectmasterii": false,
	"/customer/dashboard":        true,
	"/customer/create":           false,
	"/customer/findcustomer":     false,
}

func midelware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			if pathNoNeedLogin[r.URL.Path] {
				h.ServeHTTP(w, r)
				return
			}
			http.NotFound(w, r)
			return
		}
		dataClaims, err := tokenValidator(token.Value)
		if err != nil {
			if pathNoNeedLogin[r.RequestURI] {
				h.ServeHTTP(w, r)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    "",
				MaxAge:   -1,
				HttpOnly: true,
				Path:     "/",
			})
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		username := dataClaims.Get("username")
		values := r.URL.Query()
		values.Add("username", username.(string))
		r.URL.RawQuery = values.Encode()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func tokenValidator(tokenString string) (jwt.Claims, error) {

	token, err := jws.ParseJWT([]byte(tokenString))
	if err != nil {
		return nil, err
	}

	validator := &jwt.Validator{}
	validator.SetIssuer("FERP")

	err = token.Validate([]byte(secret), crypto.SigningMethodHS256, validator)
	return token.Claims(), err
}

func setAut(data map[string]interface{}, list []string) map[string]interface{} {
	for _, b := range list {
		data[b] = b
	}
	return data
}

func checkRoles(scompares []string, list []string) bool {
	for _, b := range scompares {
		if checkRole(b, list) {
			return true
		}
	}
	return false
}

func checkRole(scompare string, list []string) bool {
	for _, b := range list {
		if b == scompare {
			return true
		}
	}
	return false
}

func bodyToJSON(r *http.Request) map[string]string {
	body := map[string]string{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	r.Body.Close()
	return body
}

func checkAuthorization(scompares []string, user model.User, w http.ResponseWriter, r *http.Request) {
	if !checkRoles(scompares, user.Roles) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
