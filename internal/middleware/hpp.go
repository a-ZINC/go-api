package middleware

import "net/http"

type HPP struct {
	checkBody bool
	checkQuery bool
	checkBodyForContentTypeOnly string
	whiteList []string
}

func NewHPP(checkBody, checkQuery bool, checkBodyForContentTypeOnly string, allowedParameters []string) *HPP {
	return &HPP{
		checkBody: checkBody,
		checkQuery: checkQuery,
		checkBodyForContentTypeOnly: checkBodyForContentTypeOnly,
		whiteList: allowedParameters,
	}
}

func (hpp *HPP) HPP() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if hpp.checkBody && r.Method == http.MethodPost && isValidContentType(r, hpp.checkBodyForContentTypeOnly) {
				r.ParseForm()
				hpp.filterBodyParams(r, hpp.whiteList)
			}
			if hpp.checkQuery && len(r.URL.Query()) > 0 {
				hpp.filterQueryParams(r, hpp.whiteList)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func isValidContentType(r *http.Request, contentType string) bool {
	return r.Header.Get("Content-Type") == contentType
}

func (hpp *HPP) filterBodyParams(r *http.Request, whiteList []string){
	for key, val := range r.Form {
		if len(val) > 1 {
			r.Form.Set(key, val[0])
		}
		if !hpp.isValidParams(key, whiteList) {
			delete(r.Form, key)
		}
	}
}

func (hpp *HPP) filterQueryParams(r *http.Request, whiteList []string) {
	query := r.URL.Query()
	for key, val := range query {
		if len(val) > 1 {
			query.Set(key, val[0])
		}
		if !hpp.isValidParams(key, whiteList) {
			query.Del(key)
		}
	}
	r.URL.RawQuery = query.Encode()
}

func (hpp *HPP) isValidParams(key string, whiteList []string) bool {
	for _, allowedKey := range whiteList {
		if key == allowedKey {
			return true
		}
	}
	return false
}