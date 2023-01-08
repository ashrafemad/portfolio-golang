package controller

import (
	"encoding/json"
	"net/http"
	"portfolio/db"
	"portfolio/models"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project *models.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		ErrorResponse(err, w)
		return
	}
	err, project = db.SaveProject(project)
	if err != nil {
		ErrorResponse(err, w)
		return
	}
	response, _ := json.Marshal(project)
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func ListProjects(w http.ResponseWriter, r *http.Request) {
	err, projects := db.ListProjects()
	if err != nil {
		ErrorResponse(err, w)
		return
	}
	response, _ := json.Marshal(projects)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func RetrieveProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ErrorResponse(err, w)
		return
	}
	err, project := db.RetrieveProject(id)
	if err != nil {
		ErrorResponse(err, w)
		return
	}
	response, _ := json.Marshal(project)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ErrorResponse(err, w)
		return
	}
	err = db.DeleteProject(id)
	if err != nil {
		ErrorResponse(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
