package shoutouts_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jboursiquot/shoutouts/pkg/shoutouts"
)

func TestCallback(t *testing.T) {
	cases := []struct {
		scenario string
		retCode  int
	}{
		{
			scenario: "success",
			retCode:  http.StatusOK,
		},
		{
			scenario: "failure",
			retCode:  http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.scenario, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(c.retCode)
			}))
			defer ts.Close()

			cb := shoutouts.NewCallback()
			s := shoutouts.New()
			s.ResponseURL = ts.URL

			if c.retCode == 200 {
				assert.NoError(t, cb.Call(context.Background(), s))
			} else {
				assert.Error(t, cb.Call(context.Background(), s))
			}
		})
	}
}
