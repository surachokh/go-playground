package handler

import (
	"net/http"
)

func RunHandler() {
	courseItem := http.HandlerFunc(CourseHandler)
	courseList := http.HandlerFunc(CoursesHandler)
	http.Handle("/course/", enableCorsMiddleware(courseItem))
	http.Handle("/course", enableCorsMiddleware(courseList))
	http.ListenAndServe(":8080", nil)
}
