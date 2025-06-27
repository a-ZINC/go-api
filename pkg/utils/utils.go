package utils

import "net/http"

func ParseFormData(r *http.Request) string {
	name := r.Form.Get("name")
	return name
}