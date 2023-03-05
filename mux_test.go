package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/endot1231/go_todo_app/config"
)

func TestNewMux(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	cfg, err := config.New()
	ctx := context.Background()
	sut, cleanup, err := NewMux(ctx, cfg)
	defer cleanup()
	sut.ServeHTTP(w, r)
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Error("want status 200, but", resp.StatusCode)
	}

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body %v", err)
	}

	want := `{"status":"ok"}`
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
