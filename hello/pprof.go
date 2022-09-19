package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// 定义Book的结构体
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// 定义Author的结构体
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// 初始化book的切片
var books []Book

// 查询所有的Book
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

// 根据ID查询Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	log.Println(params)
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// 根据Title查询
func getBookByTitle(w http.ResponseWriter, r *http.Request) {
	var newBooks []Book
	params := mux.Vars(r) // Get params
	for _, item := range books {
		if strings.Contains(item.Title, params["title"]) {
			newBooks = append(newBooks, item)
		}
	}
	json.NewEncoder(w).Encode(newBooks)
}

// 创建Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// 修改Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	for index, item := range books {
		if item.ID == params["id"] {
			books[index] = book
			books[index].ID = item.ID
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// 删除Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	go pprof()

	// 初始化路由
	r := mux.NewRouter()

	// 增加mock数据
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Go语言", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "448744", Title: "Java语言", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "448745", Title: "java程序设计", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// 普通路由
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	// 普通路由参数
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	//r.HandleFunc("/api/books/byTitle/{title}", getBookByTitle).Methods("GET")
	// 正则路由参数，title的查询限制为英文字母,并且是小写字母，否则报：404 page not found
	r.HandleFunc("/api/books/byTitle/{title:[a-z]+}", getBookByTitle).Methods("GET")

	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	// 监听8000端口，并打出日志
	log.Fatal(http.ListenAndServe(":8000", r))
}

func pprof() error {
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		return err
	}
	return nil
}
