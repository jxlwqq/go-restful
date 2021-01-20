package healthz

import (
	"encoding/json"
	"github.com/jxlwqq/go-restful/internal/response"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	Healthz(recorder, request)
	assert.Equal(t, recorder.Code, http.StatusOK)
	res := response.Response{}
	_ = json.NewDecoder(recorder.Body).Decode(&res)
	assert.Equal(t, res.Code, http.StatusOK)
}
