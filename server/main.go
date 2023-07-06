package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	// File    *multipart.FileHeader `form:"file"`
	// File *multipart.File `form:"file"`
	// File multipart.File `form:"file"`
}

type PostId struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// File    *multipart.FileHeader `form:"file"`
	// File *multipart.File `form:"file"`
	// File multipart.File `form:"file"`
}

func checkError(err error) {
	if err != nil {
		log.Fatal("checkError: ", err)
	}
}

func getFirstParam(path string) (ps string) {
	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			ps = path[i+1:]
		}
	}
	return
}

var db *sql.DB

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryStr := "select * from Posts"

	rows, e := db.Query(queryStr)

	checkError(e)

	defer rows.Close()

	var posts []PostId

	for rows.Next() {
		var postModel PostId

		e := rows.Scan(&postModel.ID, &postModel.Title, &postModel.Content) // &postModel.File
		if e != nil {
			fmt.Println("error while row scan")
			continue
		}

		posts = append(posts, postModel)
	}

	preparedPostModel, err := json.Marshal(posts)

	checkError(err)

	w.Write(preparedPostModel)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryStr := "select * from Posts where id = $1"

	row := db.QueryRow(queryStr, getFirstParam(r.URL.Path))

	var postModel PostId

	e := row.Scan(&postModel.ID, &postModel.Title, &postModel.Content) //&postModel.File
	checkError(e)

	preparedPost, e := json.Marshal(postModel)
	checkError(e)

	w.Write([]byte(preparedPost))
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var postModel Post

	pgStr := "insert into Posts (title, content) values ($1, $2)"

	e := json.NewDecoder(r.Body).Decode(&postModel)

	// file, _, e := r.FormFile("file")
	// postModel.File = file

	checkError(e)

	_, err := db.Exec(pgStr, postModel.Title, postModel.Content) //postModel.File

	checkError(err)

	w.Write([]byte("post created successsfully"))
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().

	queryStr := "delete from Posts where id = $1"

	_, e := db.Exec(queryStr, getFirstParam(r.URL.Path))

	checkError(e)

	w.Write([]byte("deletePost success"))
}

func main() {
	e := godotenv.Load(".env")
	checkError(e)

	connStr := fmt.Sprintf("host= %s port= %s user = %s password = %s dbname = %s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, e = sql.Open("postgres", connStr)

	checkError(e)

	defer db.Close()

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/posts", getAllPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/create_post", createPost).Methods("POST")
	r.HandleFunc("/delete_post/{id}", deletePost).Methods("DELETE")

	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "127.0.0.1:8000",
	// }

	// checkError(srv.ListenAndServe())

	// checkError(http.ListenAndServe(":8080", r))

	checkError(http.ListenAndServe(":8080",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))
}
