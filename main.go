package main

import (
   "encoding/json"
   "fmt"
   "log"
   "io/ioutil"
   "net/http"
   "github.com/gorilla/mux"
)

type Article struct {
  Id string `json:"Id"`
  Title string `json:"Title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Welcome to the HomePage!")
   fmt.Println("Endpoint Hit: HomePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
   fmt.Println("Endpoint Hit: returnAllArticles")
   json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: returnSingleArticle")
   vars := mux.Vars(r)
   key := vars["id"]
   // fmt.Fprintf(w, "Key: " + key)
   for  _, article := range Articles {
      if article.Id == key {
         json.NewEncoder(w).Encode(article)
      }
   }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: createArticle")
   reqBody, _ := ioutil.ReadAll(r.Body)
   // fmt.Fprintf(w, "%+v", string(reqBody))
   var article Article
   json.Unmarshal(reqBody, &article)
   Articles = append(Articles, article)
   json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   id := vars["id"]

   for index, article := range Articles {
      if article.Id == id {
         Articles = append(append(Articles[:index], Articles[index+1:]...))
         fmt.Fprintf(w, "Item has been Deleted.")
      }
   }
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   id := vars["id"]
   reqBody, _ := ioutil.ReadAll(r.Body)
   // fmt.Fprintf(w, "%+v", string(reqBody))
   var newArticle Article
   json.Unmarshal(reqBody, &newArticle)
   for index, article := range Articles {
      if article.Id == id {
         Articles[index].Title = newArticle.Title
         Articles[index].Desc = newArticle.Desc
         Articles[index].Content = newArticle.Content
         json.NewEncoder(w).Encode(Articles[index])
      }
   }

}

func handleRequests() {
   myRouter := mux.NewRouter().StrictSlash(true)
   myRouter.HandleFunc("/", homePage)
   myRouter.HandleFunc("/articles", returnAllArticles)
   myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
   myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
   myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
   myRouter.HandleFunc("/article/{id}", returnSingleArticle)
   log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
   fmt.Println("Rest Api v2.0 - Mux Routers")
   Articles = []Article {
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
   handleRequests()
}
