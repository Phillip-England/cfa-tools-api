package net

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetBodyNew(w http.ResponseWriter, r *http.Request, v interface{}) (err error) {
	body, err := io.ReadAll(r.Body)
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}