package delivery

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFindSubstring(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/rest/substr/find", FindSubstring)

	tests := []struct {
		input          string
		expectedResult string
		expectedError  string
	}{
		{
			input:          "abcabcbb",
			expectedResult: "abc",
			expectedError:  "",
		},
		{
			input:          "bbbbb",
			expectedResult: "b",
			expectedError:  "",
		},
		{
			input:          "pwwkew",
			expectedResult: "wke",
			expectedError:  "",
		},
		{
			input:          "12345",
			expectedResult: "12345",
			expectedError:  "",
		},
		{
			input:          "",
			expectedResult: "",
			expectedError:  "требуется входная строка",
		},
		{
			input:          "aab",
			expectedResult: "ab",
			expectedError:  "",
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/rest/substr/find", strings.NewReader(test.input))
		if err != nil {
			t.Fatalf("Error creating request: %s", err)
		}

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK && rec.Code != http.StatusBadRequest {
			t.Errorf("неправильный код состояния: получено %d, ожидается %d или %d", rec.Code, http.StatusOK, http.StatusBadRequest)
		}

		if rec.Code == http.StatusOK {
			expectedJSON := `{"maxSubstring":"` + test.expectedResult + `"}`
			if rec.Body.String() != expectedJSON {
				t.Errorf("неправильный результат: получено %s, ожидается %s", rec.Body.String(), expectedJSON)
			}
		}

		if rec.Code == http.StatusBadRequest {
			expectedJSON := `{"error":"` + test.expectedError + `"}`
			if rec.Body.String() != expectedJSON {
				t.Errorf("неправильная ошибка: получено %s, ожидается %s", rec.Body.String(), expectedJSON)
			}
		}
	}
}
