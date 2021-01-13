package post

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jxlwqq/go-restful/internal/auth"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
	"net/http"
)

func RegisterHandlers(r *mux.Router, db *database.DB, logger *log.Logger, authMiddleware auth.Middleware) {
	res := resource{service: NewService(db)}
	s := r.PathPrefix("").Subrouter()
	s.HandleFunc("/posts/{id}", res.get).Methods(http.MethodGet)
	s.Use(authMiddleware.Handler)
	r.HandleFunc("/posts", res.query).Methods(http.MethodGet)
}

type resource struct {
	service Service
}

func (res resource) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	post, _ := res.service.Get(id)
	json.NewEncoder(w).Encode(post)
}

func (res resource) query(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Query")
}
