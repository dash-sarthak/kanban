package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func FormDecoder(r *http.Request, params interface{}) (*json.Decoder, error) {
	postData := make(map[string]string)

	val := reflect.ValueOf(params)

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i).Name
		if r.PostFormValue(strings.ToLower(field)) == "" {
			return &json.Decoder{}, fmt.Errorf("parameter %v cannot be empty", field)
		}
		postData[field] = r.PostFormValue(strings.ToLower(field))
	}
	jsonPostData, err := json.Marshal(postData)
	return json.NewDecoder(strings.NewReader(string(jsonPostData))), err
}
