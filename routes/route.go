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

// GetPosts godoc
// @Summary Get details of all posts
// @Description Get details of all posts
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Post
// @Router /posts [get]
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

// AddPost godoc
// @Summary Add a new post
// @Description Add a new post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param post body entity.Post true "Create post"
// @Success 200 {object} entity.Post
// @Router /posts [post]
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
