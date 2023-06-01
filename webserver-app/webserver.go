package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

/********************************************************
	model
*/

type Post struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"content"`
	CreatedAt time.Time `gorm:"created_at"`
}

type PageData struct {
	PageTitle string
	Posts     []Post
	Post      Post
}

/********************************************************
	controller
*/

// Controller holds all the variables needed for routes to perform their logic
type Controller struct {
	db *gorm.DB
}

// New creates a new instance of the routes.Controller
func NewController(db *gorm.DB) Controller {
	return Controller{
		db: db,
	}
}

func (c Controller) GetDB() *gorm.DB {
	return c.db
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ConnectToDatabase() *gorm.DB {
	dsn := "gorm:gorm@tcp(127.0.0.1:9910)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "gorm:gorm@tcp(godockerDB:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm.Open error", err)
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&Post{})
	if err != nil {
		log.Println("gorm.Open could not migrate models", err)
		panic(err)
	}

	log.Println("Successfully connected! Server is up ...")
	return db
}

func (c Controller) getPost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var post Post
	err := c.db.Where("id = ?", id).First(&post).Error
	if err != nil {
		http.Redirect(w, r, "/postNotFound", 301)
	}
	data := PageData{
		PageTitle: "Posts",
		Post:      post,
	}
	tmpl := template.Must(template.ParseFiles("templates/post.html"))
	tmpl.Execute(w, data)
}

func (c Controller) getPosts(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	c.db.Find(&posts)
	data := PageData{
		PageTitle: "My posts",
		Posts:     posts,
	}
	tmpl := template.Must(template.ParseFiles("templates/posts.html"))
	tmpl.Execute(w, data)
}

func (c Controller) createPost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl := template.Must(template.ParseFiles("templates/create_post.html"))
		data := PageData{
			PageTitle: "Create Post",
		}
		tmpl.Execute(w, data)
	case http.MethodPost:
		var post Post
		post.Content = r.FormValue("content")
		err := c.db.Create(&post).Error
		if err != nil {
			http.Redirect(w, r, "/couldNotCreate", 301)
		}
		http.Redirect(w, r, "/post/"+strconv.Itoa(post.Id), 301)
	default:
		http.Redirect(w, r, "/methodNotSupported", 301)
	}
}

func (c Controller) status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status, %s!", r.URL.Path[1:])
}

/*
https://gowebexamples.com/templates/
*/
func main() {
	r := mux.NewRouter()

	// We connect to the database using the configuration generated from the environment variables.
	db := ConnectToDatabase()

	// A new instance of controller is created
	controller := NewController(db)

	r.Use(mux.CORSMethodMiddleware(r))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/posts", controller.getPosts).Methods(http.MethodGet)
	r.HandleFunc("/post/{id:[0-9]+}", controller.getPost).Methods(http.MethodGet)
	r.HandleFunc("/post_create", controller.createPost).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc ("/status", controller.status).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}

