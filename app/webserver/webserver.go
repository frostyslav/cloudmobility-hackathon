package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frostyslav/cloudmobility-hackathon/app/image"
	"github.com/frostyslav/cloudmobility-hackathon/app/render"
	"github.com/gorilla/mux"
)

var Serve http.Handler

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/", status).Methods("GET")
	r.HandleFunc("/func_create", funcCreate).Methods("POST")
	//r.HandleFunc("/func_status", funcStatus).Methods("GET")

	Serve = r
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK\n")
}

func funcCreate(w http.ResponseWriter, r *http.Request) {
	req := &CreateFuncFromRepoRequest{}
	err := render.ReadJSON(r, req)
	if err != nil {
		log.Print(err)
	}

	log.Printf("repo url: %s, repo tag: %s, directory: %s, lang: %s", req.Repo.URL, req.Repo.Tag, req.Repo.Path, req.Language)
	image, _ := image.Build()
	resp := &ImageURLResponse{
		URL: image,
	}

	render.WriteJSONwithCode(w, resp, 200)
}

// ImageURLResponse - response for /create_func_from_repo.
// swagger:model ImageURLResponse
type ImageURLResponse struct {
	// Image URL
	URL string `json:"image_url"`
}

type CreateFuncFromRepoRequest struct {
	Repo     Repo   `json:"repo"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

type Repo struct {
	URL  string `json:"url"`
	Tag  string `json:"tag"`
	Path string `json:"path"`
}
