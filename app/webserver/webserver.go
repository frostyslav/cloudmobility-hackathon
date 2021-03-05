package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frostyslav/cloudmobility-hackathon/app/controller"
	"github.com/frostyslav/cloudmobility-hackathon/app/model"
	"github.com/frostyslav/cloudmobility-hackathon/app/render"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var Serve http.Handler
var hashmap *model.Hash

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/", status).Methods("GET")
	r.HandleFunc("/func_create", funcCreate).Methods("POST", "OPTIONS")
	r.HandleFunc("/func_status/{id}", funcStatus).Methods("GET", "OPTIONS")

	m := model.Hash{}
	hashmap = m.New()

	Serve = r
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK\n")
}

func funcCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if (*r).Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	req := &model.FuncCreateRequest{}
	err := render.ReadJSON(r, req)
	if err != nil {
		log.Print(err)
	}

	if req.Repo.URL == "" && req.Code == "" {
		render.WriteJSONwithCode(w, nil, 400)
		return
	}

	myuuid := uuid.New()

	if req.Repo.URL != "" {
		go controller.RunFromRepo(hashmap, req.Repo.URL, req.Repo.Tag, req.Repo.Path, myuuid.String())
	} else if req.Code != "" {
		if req.Language == "" {
			req.Language = "go"
		}
		go controller.RunFromCode(hashmap, req.Code, myuuid.String(), req.Language)
	} else {
		render.WriteJSONwithCode(w, nil, 400)
		return
	}

	hashmap.SetStatus(myuuid.String(), "")

	resp := &model.FuncCreateResponse{
		ID: myuuid.String(),
	}

	render.WriteJSONwithCode(w, resp, 200)
}

func funcStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if (*r).Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	vars := mux.Vars(r)

	resp := &model.FuncStatusResponse{
		ID:     vars["id"],
		Status: hashmap.GetStatus(vars["id"]),
	}

	render.WriteJSONwithCode(w, resp, 200)
}
