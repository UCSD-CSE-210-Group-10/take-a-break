package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Config struct {
	ClientID        string
	ClientSecret    string
	AuthURL         string
	TokenURL        string
	RedirectURL     string
	ClientURL       string
	TokenSecret     string
	TokenExpiration int64
	PostURL         string
}

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

var config = Config{
	ClientID:        os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret:    os.Getenv("GOOGLE_CLIENT_SECRET"),
	AuthURL:         "https://accounts.google.com/o/oauth2/v2/auth",
	TokenURL:        "https://oauth2.googleapis.com/token",
	RedirectURL:     os.Getenv("REDIRECT_URL"),
	ClientURL:       os.Getenv("CLIENT_URL"),
	TokenSecret:     os.Getenv("TOKEN_SECRET"),
	TokenExpiration: 36000,
	PostURL:         "https://jsonplaceholder.typicode.com/posts",
}

func AuthParams() string {
	params := url.Values{}
	params.Set("client_id", config.ClientID)
	params.Set("redirect_uri", config.RedirectURL)
	params.Set("response_type", "code")
	params.Set("scope", "openid profile email")
	params.Set("access_type", "offline")
	params.Set("state", "standard_oauth")
	params.Set("prompt", "consent")
	return params.Encode()
}

func GetTokenParams(code string) string {
	params := url.Values{}
	params.Set("client_id", config.ClientID)
	params.Set("client_secret", config.ClientSecret)
	params.Set("code", code)
	params.Set("grant_type", "authorization_code")
	params.Set("redirect_uri", config.RedirectURL)
	return params.Encode()
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := tokenCookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.TokenSecret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetAuthURLHandler(w http.ResponseWriter, r *http.Request) {
	authURL := fmt.Sprintf("%s?%s", config.AuthURL, AuthParams())
	json.NewEncoder(w).Encode(map[string]string{"url": authURL})
}

func GetAuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Authorization code must be provided", http.StatusBadRequest)
		return
	}
	tokenParams := GetTokenParams(code)
	resp, err := http.Post(config.TokenURL, "application/x-www-form-urlencoded", strings.NewReader(tokenParams))
	if err != nil {
		http.Error(w, "Failed to exchange authorization code for token", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	var tokenResp struct {
		IDToken string `json:"id_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		http.Error(w, "Failed to parse token response", http.StatusInternalServerError)
		return
	}
	if tokenResp.IDToken == "" {
		http.Error(w, "Auth error", http.StatusBadRequest)
		return
	}
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenResp.IDToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenSecret), nil
	})
	if err != nil {
		http.Error(w, "Failed to decode token", http.StatusInternalServerError)
		return
	}
	user := User{
		Name:    claims["name"].(string),
		Email:   claims["email"].(string),
		Picture: claims["picture"].(string),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Second * time.Duration(config.TokenExpiration)).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.TokenSecret))
	if err != nil {
		http.Error(w, "Failed to sign token", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Second * time.Duration(config.TokenExpiration)),
		HttpOnly: true,
	})
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user": user,
	})
}

func LoggedInHandler(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]bool{"loggedIn": false})
		return
	}
	tokenString := tokenCookie.Value
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenSecret), nil
	})
	if err != nil {
		json.NewEncoder(w).Encode(map[string]bool{"loggedIn": false})
		return
	}
	user := claims["user"].(map[string]interface{})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Second * time.Duration(config.TokenExpiration)).Unix(),
	})
	tokenString, err = token.SignedString([]byte(config.TokenSecret))
	if err != nil {
		http.Error(w, "Failed to sign token", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Second * time.Duration(config.TokenExpiration)),
		HttpOnly: true,
	})
	json.NewEncoder(w).Encode(map[string]interface{}{
		"loggedIn": true,
		"user":     user,
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	})
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out"})
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
}

func GetPostsHandlerWithAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenCookie, err := r.Cookie("token")
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        tokenString := tokenCookie.Value
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(config.TokenSecret), nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        handler(w, r)
    }
}
