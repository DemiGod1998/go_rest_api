package routes

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json: "id"`
	Title string `json: "title"`
	Text  string `json: "text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text1"}}
}

func GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "Error marshalling the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("content-type", "application/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "Error unmarshalling the request"}`))
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)
	resp.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(posts)
	resp.Write(result)

}
