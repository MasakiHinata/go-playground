package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gorilla/sessions"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", MyHandler)
	http.ListenAndServe(":8080", nil)
}

// NewCookieStore
// Cookieを用いたSessionStoreを取得
// 生成時に秘密鍵を渡す
// 実際にはStoreKeyは環境変数から取得し、コードには書かない
// crypto/rand or securecookie.GenerateRandomKey(32) などを利用してランダムなキーを生成する
var store = sessions.NewCookieStore([]byte("secret"))
var tokens []string

func init() {
	store.Options.HttpOnly = true
	store.Options.SameSite = http.SameSiteDefaultMode
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	// 存在するセッションを取得、または新規作成
	session, _ := store.Get(r, "session-name")

	if count, ok := session.Values["count"]; ok {
		session.Values["count"] = count.(int) + 1
	} else {
		session.Values["count"] = 1
	}
	log.Println(session.Values["count"])

	if token, ok := session.Values["token"]; ok {
		t := token.(string)
		if contain(tokens, t) {
			// トークンが一致した場合は、トークンを更新
			log.Println("valid token")

			// 現在のトークンを削除
			removeValue(&tokens, t)
			// トークンを作成
			newToken := createToken()
			tokens = append(tokens, newToken)
			session.Values["token"] = newToken
		} else {
			// トークンが一致しない、新しいトークンを作成
			log.Println("invalid token")

			// トークンを作成
			newToken := createToken()
			tokens = append(tokens, newToken)
			session.Values["token"] = newToken
		}
	} else {
		// トークンが設定されていない場合は、トークンを設定
		log.Println("Create new Token")
		token := createToken()
		tokens = append(tokens, token)
		session.Values["token"] = token
	}

	log.Println(tokens)

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func contain(values []string, value string) bool {
	for _, v := range values {
		if value == v {
			return true
		}
	}
	return false
}

func removeValue(values *[]string, value string)  {
	for i, v := range *values {
		if value == v {
			s := *values
			*values = append(s[:i], s[i+1:]...)
		}
	}
}

func createToken() string {
	h := md5.New()
	salt := "golang%^7&8888"
	io.WriteString(h, salt+time.Now().String())
	return fmt.Sprintf("%x", h.Sum(nil))
}

//func Login(w http.ResponseWriter, r *http.Request)  {
//	var maxLifetime int64 = 60
//	session, _ := store.Get(r, "SID")
//
//	if createTime, ok := session.Values["createtime"]; ok {
//		if createTime.(int64) + maxLifetime < time.Now().Unix(){
//			session.Values["createtime"]
//		}
//	}else {
//		session.Values["createtime"] = time.Now().Unix()
//	}
//}
