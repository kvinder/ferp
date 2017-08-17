package app

import (
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

	//Admin
	midelwareMux.HandleFunc("/admin/create", adminCreate)
	midelwareMux.HandleFunc("/admin/list", adminList)
	midelwareMux.HandleFunc("/admin/user", adminUser)
	midelwareMux.HandleFunc("/admin/user/update", adminUpdate)

	//FAMS
	midelwareMux.HandleFunc("/fams/request", famsRequest)

	mux.Handle("/", midelware(midelwareMux))

}

var pathNoNeedLogin = map[string]bool{
	"/":                  true,
	"/login":             true,
	"/fa/dashboard":      true,
	"/admin/create":      false,
	"/admin/list":        false,
	"/admin/user":        false,
	"/admin/user/update": false,
	"/fams/request":      false,
}

func midelware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			if pathNoNeedLogin[r.RequestURI] {
				h.ServeHTTP(w, r)
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
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
