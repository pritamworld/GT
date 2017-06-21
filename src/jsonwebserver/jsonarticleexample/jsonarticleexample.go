package main
//mux means "HTTP request Multiplexer"
import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Articles []Article 

//Home or default page
//http://localhost:10000
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}


//Return/Display 1 article
//http://localhost:10000/article/1/1/2/
func returnOneArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["key"]
    var1 := vars["var1"]
    var2 := vars["var2"]

    fmt.Println("Var 1: " + var1)
    fmt.Println("Var 2: " + var2)
    fmt.Fprintf(w, "Key: " + key)
}

//Return all articles
//http://localhost:10000/all
func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "All Articles")
    articles := Articles{
        Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }  
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(articles)
}
//http://localhost:10000/add
func addArticle(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Adds an article to list of articles")
    fmt.Println("Endpoint Hit: addArticle")
}

//http://localhost:10000/delete
func delArticle(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "deletes a specific article")
    fmt.Println("Endpoint Hit: delArticle")
}

/*
func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/all", returnAllArticles)
    http.HandleFunc("/single", returnArticle)
    http.HandleFunc("/delete", delArticle)
    http.HandleFunc("/add", addArticle)
    log.Fatal(http.ListenAndServe(":8081", nil))
}
*/

func handleRequests() {

    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage) //http://localhost:10000
    myRouter.HandleFunc("/all", returnAllArticles) //http://localhost:10000/all
    //http://localhost:10000/article/1/1/2/
    myRouter.HandleFunc("/article/{key}/{var1}/{var2}/", returnOneArticle)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}