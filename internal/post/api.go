package post

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jxlwqq/go-restful/internal/auth"
	"github.com/jxlwqq/go-restful/internal/response"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
	"net/http"
)

func RegisterHandlers(r *mux.Router, db *database.DB, logger *log.Logger, authMiddleware auth.Middleware) {
	res := resource{service: NewService(db)}
	s := r.PathPrefix("").Subrouter()
	s.HandleFunc("/posts", res.Create).Methods(http.MethodPost)
	s.Use(authMiddleware.Handler)
	r.HandleFunc("/posts", res.Query).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", res.Get).Methods(http.MethodGet)
}

type resource struct {
	service Service
}

func (res resource) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	post, _ := res.service.Get(id)
	go func() {
		_ = res.service.IncrementViewCount(post)
	}()
	response.Write(w, post, http.StatusOK)
}

func (res resource) Query(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Query")
}

func (res resource) Create(w http.ResponseWriter, r *http.Request) {
	req := CreateRequest{}
	_ = json.NewDecoder(r.Body).Decode(&req)
	post, _ := res.service.Create(req)
	response.Write(w, post, http.StatusCreated)
}
