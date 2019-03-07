package cb

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kisamoto/janus/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(*testing.T, *http.Request, *httptest.ResponseRecorder)
	}{
		{
			scenario: "with wrong predicate given",
			function: testWrongPredicate,
		},
		{
			scenario: "when the upstream respond successfully",
			function: testSuccessfulUpstreamRetry,
		},
		{
			scenario: "when the upstream fails to respond",
			function: testFailedUpstreamRetry,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			test.function(t, r, w)
		})
	}
}

func testWrongPredicate(t *testing.T, r *http.Request, w *httptest.ResponseRecorder) {
	cfg := Config{
		Name:      "example",
		Predicate: "this is wrong",
	}
	mw := NewCBMiddleware(cfg)

	mw(http.HandlerFunc(test.Ping)).ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func testSuccessfulUpstreamRetry(t *testing.T, r *http.Request, w *httptest.ResponseRecorder) {
	mw := NewCBMiddleware(Config{Name: "example"})

	mw(http.HandlerFunc(test.Ping)).ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func testFailedUpstreamRetry(t *testing.T, r *http.Request, w *httptest.ResponseRecorder) {
	mw := NewCBMiddleware(Config{Name: "example"})

	mw(test.FailWith(http.StatusBadGateway)).ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadGateway, w.Code)
}
