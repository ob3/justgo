package justgo

import (
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/not_found", nil)
	require.NoError(t, err, "failed to create a request")

	pingHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
