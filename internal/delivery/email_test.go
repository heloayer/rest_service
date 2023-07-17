package delivery

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckEmail(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader("test@example.com"))

	w := httptest.NewRecorder()

	router := gin.Default()

	router.POST("/test", CheckEmail)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}
	expected := `{"emails":["test@example.com"]}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, w.Body.String())
	}
}
