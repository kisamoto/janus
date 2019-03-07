package web

import (
	"testing"

	"net/http"

	"github.com/kisamoto/janus/pkg/api"
	"github.com/kisamoto/janus/pkg/router"
	"github.com/kisamoto/janus/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	r := router.NewChiRouter()
	r.GET("/status", NewOverviewHandler(&api.Configuration{}))

	ts := test.NewServer(r)
	defer ts.Close()

	res, _ := ts.Do(http.MethodGet, "/status", make(map[string]string))
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
}
