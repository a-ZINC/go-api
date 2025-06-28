package router

import "net/http"

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/teachers", teacherHandler)
	mux.HandleFunc("/students", studentHandler)
	mux.HandleFunc("/exams", examHandler)
	return mux
}
