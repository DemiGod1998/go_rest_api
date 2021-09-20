package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"go_rest_api/entity"
	"go_rest_api/repository"
)

var (
	repo repository.PostRepository = repository.NewPostrepository()
)

func GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "Error getting the posts"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "Error unmarshalling the request"}`))
		return
	}

	post.ID = rand.Int63()
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
