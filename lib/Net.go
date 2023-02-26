package lib

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetBody(w http.ResponseWriter, r *http.Request, v interface{}) (err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}

func GetURLParam(path string) (param string) {
	parts := strings.Split(path, "/")
	param = parts[len(parts)-1]
	return param
}

func HttpCookie(w http.ResponseWriter, name string, value string, expireMinutes int) {
	minutes := time.Duration(expireMinutes)
	expire := time.Now().Add(minutes * time.Minute)
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expire,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}

func Serve() {
	log.Println("Serving application on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
