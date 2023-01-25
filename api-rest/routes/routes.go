package routes

import (
	"log"
	"net/http"

	"api-rest-go/controllers"
	"api-rest-go/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Index)
	r.HandleFunc("/api/posts", controllers.GetAllPosts).Methods("Get")
	r.HandleFunc("/api/posts", controllers.CreateNewPost).Methods("Post")
	r.HandleFunc("/api/posts/{id}", controllers.DeletePost).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
