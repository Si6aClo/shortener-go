package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"net/http/httptest"
	r "shortener/router"
	"strings"
	"testing"
)

func TestMakeShorterHandler(t *testing.T) {
	if err := godotenv.Load(); err != nil {
	}
	router := gin.Default()
	r.SetupRouter(router)
	w := httptest.NewRecorder()
	//pgContext, err := db.NewDB()
	//if err != nil {
	//	t.Fatal(err)
	//}

	t.Run("Check short link create", func(t *testing.T) {
		// send test query on make_shorter with long_url in body and check response
		req, err := http.NewRequest("POST", "/api/v1/make_shorter", strings.NewReader(`{"url": "https://google.com"}`))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Code)
		}
		fmt.Println(w.Body.String())
		// check that short link was created in db
		//shortLink, err := tx.Query("SELECT short_url FROM url_storage WHERE long_url = $1", "https://google.com")
		//if err != nil {
		//	t.Fatal(err)
		//}
		//if !shortLink.Next() {
		//	t.Errorf("Expected short link, got nothing")
		//}
	})
}
