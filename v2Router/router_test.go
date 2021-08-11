package v2Router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
)

func TestUserSave(t *testing.T) {
	username := "lisi"
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "User: "+username+" saved already!", w.Body.String())
}
