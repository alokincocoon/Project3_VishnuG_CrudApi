package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/edgedb/edgedb-go"
	"github.com/gorilla/mux"
	"github.com/vishnu-g-k/student-management/models"
	utils "github.com/vishnu-g-k/student-management/utils"
)

type JsonResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func respondWithJSON(r *http.Request, w http.ResponseWriter, code int, payload interface{}, status string) {
	var jsonResponse JsonResponse
	jsonResponse.Status = status
	jsonResponse.Data = payload
	response, _ := json.Marshal(jsonResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

	logger_ := utils.Logger{
		Method:     r.Method,
		Path:       r.URL.Path,
		Response:   jsonResponse,
		StatusCode: code,
		Status:     status,
	}
	go utils.Log(&logger_)
}

func respondWithError(r *http.Request, w http.ResponseWriter, code int, message string) {
	respondWithJSON(r, w, code, map[string]string{"message": message}, "error")
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	respondWithError(r, w, http.StatusNotFound, "Page Not Found")
}

func GetStudents(dbClient *edgedb.Client, w http.ResponseWriter, r *http.Request) {
	students, err := models.GetStudents(dbClient)
	if err != nil {
		respondWithError(r, w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(r, w, http.StatusOK, students, "success")
}

func CreateStudent(dbClient *edgedb.Client, w http.ResponseWriter, r *http.Request) {
	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		log.Println("Invalid request payload")
		respondWithError(r, w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := utils.ValidateStudent(&student); err != nil {
		log.Println("Validation Failed")
		respondWithError(r, w, http.StatusBadRequest, err.Error())
		return
	}

	insertedId, err := models.CreateStudent(dbClient, &student)
	if err != nil {
		respondWithError(r, w, http.StatusInternalServerError, err.Error())
		return
	}
	student.ID = insertedId
	respondWithJSON(r, w, http.StatusCreated, student, "success")
}

func GetStudentDetails(dbClient *edgedb.Client, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["id"]
	uuid, err := edgedb.ParseUUID(studentId)
	if err != nil {
		respondWithError(r, w, http.StatusBadRequest, "Invalid ID")
		return
	}
	student, err := models.GetStudentDetails(dbClient, uuid)
	if err != nil {
		respondWithError(r, w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(r, w, http.StatusOK, student, "success")
}

func UpdateStudentDetails(dbClient *edgedb.Client, w http.ResponseWriter, r *http.Request) {
	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		log.Println("Invalid request payload")
		respondWithError(r, w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	studentId := vars["id"]
	uuid, er := edgedb.ParseUUID(studentId)
	if er != nil {
		respondWithError(r, w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := utils.ValidateStudent(&student); err != nil {
		log.Println("Validation Failed")
		respondWithError(r, w, http.StatusBadRequest, err.Error())
		return
	}

	student.ID = uuid

	err := models.UpdateStudentDetails(dbClient, &student)
	if err != nil {
		respondWithError(r, w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(r, w, http.StatusOK, student, "success")
}

func DeleteStudentDetails(dbClient *edgedb.Client, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["id"]
	uuid, err := edgedb.ParseUUID(studentId)
	if err != nil {
		respondWithError(r, w, http.StatusBadRequest, "Invalid ID")
		return
	}
	student, er := models.DeleteStudent(dbClient, uuid)
	fmt.Println("Student Deleted: ", student)
	if er != nil {
		respondWithError(r, w, http.StatusInternalServerError, er.Error())
		return
	}
	respondWithJSON(r, w, http.StatusOK, map[string]string{"message": fmt.Sprintf("Student with ID '%s' deleted successfully", uuid)}, "success")
}
