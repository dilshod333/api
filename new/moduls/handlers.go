package moduls

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret_key")

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
    var credent Credentials
    err := json.NewDecoder(r.Body).Decode(&credent)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    expectedPassword, ok := users[credent.Username]
    if !ok || expectedPassword != credent.Password {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(time.Minute * 1)
    claims := &Claims{
        Username: credent.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(secretKey)

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
}

func Home(w http.ResponseWriter, r *http.Request) {
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

    tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

    if err != nil || !tkn.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    w.Write([]byte(fmt.Sprintf("Hello %s", claims.Username)))
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

    tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

    if err != nil || !tkn.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    fmt.Fprintf(w, "Hello Dilshodbek...")
}
