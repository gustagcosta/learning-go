package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api-rest-go/database"
	"api-rest-go/models"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var p []models.Post
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post
	json.NewDecoder(r.Body).Decode(&newPost)
	database.DB.Create(&newPost)
	json.NewEncoder(w).Encode(newPost)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var post models.Post
	database.DB.Delete(&post, id)
	json.NewEncoder(w).Encode(post)
}
