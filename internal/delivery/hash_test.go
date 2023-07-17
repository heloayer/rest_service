package delivery

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCalcHash(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader("test data"))

	w := httptest.NewRecorder()

	router := gin.Default()

	router.POST("/test", CalcHash)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Errorf("Expected status %d but got %d", http.StatusAccepted, w.Code)
	}
}

func TestGetResult(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/test/e2a0d4e3f3a8e1e7a3a4f3a8e1e7a3a4f3a8e1e7a3a4f3a8e1e7a3a4f3a8e1e7", nil)

	w := httptest.NewRecorder()

	router := gin.Default()

	router.GET("/test/:request_id", GetResult)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

}
