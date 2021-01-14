package healthz

import (
	"github.com/gorilla/mux"
	"github.com/jxlwqq/go-restful/internal/response"
	"net/http"
)

func RegisterHandlers(r *mux.Router)  {
	r.HandleFunc("/healthz", Healthz)
}

func Healthz(w http.ResponseWriter, r *http.Request)  {
	response.New(w, nil, http.StatusOK)
}
