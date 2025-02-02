package auth

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Add("Content-Type", "text/html; charset=UTF-8")
		fp := path.Join("templates", "login.html")

		templ, err := template.ParseFiles(fp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := templ.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		fmt.Println("POST method Invoked")
		var credentials Credentials
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		// Access form values
		username := r.FormValue("username")
		password := r.FormValue("password")
		credentials.Username = username
		credentials.Password = password
		expectedPassword, ok := users[credentials.Username]
		fmt.Println(expectedPassword)
		if !ok || expectedPassword != credentials.Password {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println(username, password)

			return
		}

		expirationTime := time.Now().Add(time.Minute * 5)

		claims := &Claims{
			Username: credentials.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w,
			&http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func IsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	isAuth := true
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		// w.WriteHeader(http.StatusBadRequest)
		return false
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		// w.WriteHeader(http.StatusBadRequest)
		return false
	}

	if !tkn.Valid {
		// w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return isAuth

}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Clear the token cookie by setting its expiration time to the past
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Now().Add(-time.Hour),
		})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
