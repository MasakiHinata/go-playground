package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
	redirectUrl       = "http://localhost:8080/oauth/google/callback"
)

var (
	config *oauth2.Config
)

func main() {
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if googleClientID == "" {
		log.Println("GOOGLE_CLIENT_ID is not given.")
		return
	}
	if googleClientSecret == "" {
		log.Println("GOOGLE_CLIENT_SECRET is not given")
		return
	}

	config = &oauth2.Config{
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"openid", "email", "profile"},
		RedirectURL: redirectUrl,
	}

	http.HandleFunc("/", Index)
	http.HandleFunc("/oauth/google", Oauth)
	http.HandleFunc("/oauth/google/callback", Callback)
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/template/index.html")
	t.Execute(w, nil)
}

func Oauth(w http.ResponseWriter, r *http.Request) {
	url := config.AuthCodeURL("0000")
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//tokenSource := config.TokenSource(r.Context(), token)
	//client := oauth2.NewClient(r.Context(), tokenSource)

	userInfoUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", url.QueryEscape(token.AccessToken))
	resp, _ := http.Get(userInfoUrl)
	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)
	w.Write(contents)
}
