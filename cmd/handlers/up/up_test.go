package up

import (
	"go-store/cmd/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func init() {
	r = router.SetupRouter()
	AddUpV1(r)
}

func TestGetUp(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/up", nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}
