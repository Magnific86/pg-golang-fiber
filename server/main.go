package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func checkError(err error) {
	if err != nil {
		log.Fatal("checkError: ", err)
	}
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("success"))
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("success"))
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("success"))
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("success"))
}

func main() {
	e := godotenv.Load(".env")
	checkError(e)

	connStr := fmt.Sprintf("host= %s port= %s user = %s password = %s dbname = %s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, e := sql.Open("postgres", connStr)
	checkError(e)
	defer db.Close()

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/posts", getAllPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/create_post", createPost).Methods("POST")
	r.HandleFunc("/delete_post/{id}", createPost).Methods("DELETE")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}

	checkError(srv.ListenAndServe())

	http.ListenAndServe(":8080", r)

	fmt.Println("fmsdrgbid")
}
