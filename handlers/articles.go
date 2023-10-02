package handlers

import (
	"articlesfeedapi/dal"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func GetLatestArticlesHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: Add team support
	langStr := r.URL.Query().Get("lang")
	if langStr == "" {
		langStr = "1"
	}
	langID, err := strconv.Atoi(langStr)
	if err != nil {
		http.Error(w, "Invalid Page Size", http.StatusBadRequest)
		return
	}

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	currentPage, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid Page Number", http.StatusBadRequest)
		return
	}

	sizeStr := r.URL.Query().Get("size")
	if sizeStr == "" {
		sizeStr = "10"
	}
	resultsPerPage, err := strconv.Atoi(sizeStr)
	if err != nil {
		http.Error(w, "Invalid Page Size", http.StatusBadRequest)
		return
	}
	if resultsPerPage > 50 {
		resultsPerPage = 50
	}

	sourceIDs := []int{}
	sourceStr := r.URL.Query().Get("sources")
	if sourceStr != "" {
		sourceIDs, err = stringToIntSlice(sourceStr)
		if err != nil {
			log.Fatalf("Error converting string to int slice: %v", err)
		}
	}

	fmt.Println(time.Now().Format(time.RFC3339) + " - Trigger GetLatestArticlesHandler")
	fmt.Println("LanguageID =" + langStr)
	fmt.Println("Page =" + pageStr)
	fmt.Println("Size =" + sizeStr)
	fmt.Println("SourceIDs =" + sourceStr)
	fmt.Println("Client IP =" + r.RemoteAddr)

	articles, err := dal.GetArticles(
		sourceIDs,
		langID,
		resultsPerPage,
		currentPage)

	// Marshal the articles slice into JSON
	articlesJSON, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, "Failed to marshal articles into JSON", http.StatusInternalServerError)
		return
	}

	// Set appropriate content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(articlesJSON)
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	langStr := r.URL.Query().Get("lang")
	if langStr == "" {
		langStr = "1"
	}
	langID, err := strconv.Atoi(langStr)
	if err != nil {
		http.Error(w, "Invalid Page Size", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["article_id"]

	articleID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	fmt.Println(time.Now().Format(time.RFC3339) + " - Trigger GetArticleHandler")
	fmt.Println("LanguageID =" + langStr)
	fmt.Println("ArticleID =" + idStr)
	fmt.Println("Client IP =" + r.RemoteAddr)

	article, err := dal.GetArticle(
		articleID,
		langID)

	// Marshal the article into JSON
	articleJSON, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Failed to marshal articles into JSON", http.StatusInternalServerError)
		return
	}

	// Set appropriate content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(articleJSON)
}

func stringToIntSlice(sourceStr string) ([]int, error) {
	strs := strings.Split(sourceStr, ",")
	ints := make([]int, len(strs))
	for i, str := range strs {
		var err error
		ints[i], err = strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}
