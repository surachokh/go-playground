package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

var CourseList []Course

func init() {
	CourseJSON := `[
		{
			"id":1,
			"name":"Python",
			"price":2590,
			"instructor":"Care"
		},
		{
			"id":2,
			"name":"Java",
			"price":1690,
			"instructor":"Care"
		},
		{
			"id":3,
			"name":"HTML",
			"price":3590,
			"instructor":"Care"
		}
	]`

	err := json.Unmarshal([]byte(CourseJSON), &CourseList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestId := -1
	for _, course := range CourseList {
		if highestId < course.Id {
			highestId = course.Id
		}
	}
	return highestId + 1
}

func CoursesHandler(w http.ResponseWriter, r *http.Request) {
	CourseJSON, err := json.Marshal(CourseList)

	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if CourseJSON != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(CourseJSON)
		}
	case http.MethodPost:
		var newCourse Course
		BodyByte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		errMarshal := json.Unmarshal(BodyByte, &newCourse)
		if errMarshal != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newCourse.Id = getNextID()
		CourseList = append(CourseList, newCourse)
		newCourseJSON, err := json.Marshal(CourseList)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(newCourseJSON)
	}
}

func findID(ID int) (int, error) {
	for i, course := range CourseList {
		if course.Id == ID {
			return i, nil
		}
	}
	err := "Cannot find course with ID : " + strconv.Itoa(ID)
	return 0, errors.New(err)
}

func CourseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "course/")
	ID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	index, errID := findID(ID)

	if errID != nil {
		http.Error(w, errID.Error(), http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		CourseJSON, err := json.Marshal(CourseList[index])
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(CourseJSON)
	case http.MethodPut:
		var newCourse Course
		BodyByte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		errMarshal := json.Unmarshal(BodyByte, &newCourse)
		if errMarshal != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for index, course := range CourseList {
			if course.Id == ID {
				CourseList[index] = newCourse
				newCourseJSON, err := json.Marshal(CourseList)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write(newCourseJSON)
				return
			}
		}
		http.Error(w, fmt.Sprintf("No Course that match with ID : %d", ID), http.StatusNotFound)

	}

}
