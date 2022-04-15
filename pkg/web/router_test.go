package web

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	r := NewRouter()
	called := false
	// Ensure that routes exist on the mux router
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		called = true
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, called)
}
