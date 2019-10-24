package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/http/rest"
	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetTransaction(t *testing.T) {
	refreshUserTable()
	seedOneUser()
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("this is the error: %v", err)
	}
	rr := httptest.NewRecorder()
	result := []models.User{}
	handler := rest.Handler(*lister)
	handler.ServeHTTP(rr, req)
	strResult := rr.Body.String()
	err = json.Unmarshal([]byte(strResult), &result)
	assert.Equal(t, 200, rr.Code)
	fmt.Println(result)
	assert.Equal(t, 1, len(result))

}
