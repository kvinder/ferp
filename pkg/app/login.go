package app

import (
	"ferp/pkg/model"
	"ferp/pkg/view"
	"log"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		_, err := model.UserOnLogin(r)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		view.Login(w, nil)
		return
	}
	username := r.FormValue("form-field-username")
	password := r.FormValue("form-field-password")
	token, err := createToken(username, password)
	if err != nil {
		data := map[string]interface{}{
			"loginFail": "Invalid username and password!",
		}
		view.Login(w, data)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   int(24 * 60 * time.Minute / time.Second),
		HttpOnly: true,
		Path:     "/",
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Path:     "/",
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

const secret = "foamtecintl"

func createToken(username, password string) (string, error) {
	user := model.GetByUsername(username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	claims := jws.Claims{}
	claims.Set("username", username)
	claims.SetIssuer("FERP")
	now := time.Now()
	claims.SetIssuedAt(now)
	claims.SetExpiration(now.AddDate(0, 0, 3))
	tokenStruct := jws.NewJWT(claims, crypto.SigningMethodHS256)

	serialized, err := tokenStruct.Serialize([]byte(secret))
	if err != nil {
		log.Fatal("error : ", err.Error())
	}

	token := string(serialized)
	return token, nil
}
