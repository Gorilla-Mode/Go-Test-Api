package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func fixtureObstacles() []Obstacle {
	return []Obstacle{
		{ID: 1, Name: "Obstacle 1", Geometry: "Polygon((0 0, 10 0, 10 10, 0 10))"},
		{ID: 2, Name: "Obstacle 2", Geometry: "Polygon((10 0, 20 0, 20 10, 10 10))"},
	}
}

func resetObstacles(t *testing.T) {
	t.Helper()

	original := obstacles
	obstacles = fixtureObstacles()
	t.Cleanup(func() {
		obstacles = original
	})
}

func performRequest(router http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, path, body)
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	return recorder
}

func TestGetObstacles(t *testing.T) {
	resetObstacles(t)
	router := setupRouter()

	response := performRequest(router, http.MethodGet, "/obstacles", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.JSONEq(t, `[
		{"id":1,"name":"Obstacle 1","geometry":"Polygon((0 0, 10 0, 10 10, 0 10))"},
		{"id":2,"name":"Obstacle 2","geometry":"Polygon((10 0, 20 0, 20 10, 10 10))"}
	]`, response.Body.String())
}

func TestGetObstacle(t *testing.T) {
	router := setupRouter()
	tests := []struct {
		name     string
		path     string
		wantCode int
		wantBody string
	}{
		{
			name:     "existing obstacle",
			path:     "/obstacles/1",
			wantCode: http.StatusOK,
			wantBody: `{"id":1,"name":"Obstacle 1","geometry":"Polygon((0 0, 10 0, 10 10, 0 10))"}`,
		},
		{
			name:     "invalid ID",
			path:     "/obstacles/not-a-number",
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"invalid obstacle ID"}`,
		},
		{
			name:     "missing obstacle",
			path:     "/obstacles/3",
			wantCode: http.StatusNotFound,
			wantBody: `{"error":"obstacle not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetObstacles(t)

			response := performRequest(router, http.MethodGet, tt.path, nil)

			assert.Equal(t, tt.wantCode, response.Code)
			assert.JSONEq(t, tt.wantBody, response.Body.String())
		})
	}
}

func TestPostObstacle(t *testing.T) {
	router := setupRouter()
	newObstacle := Obstacle{
		ID:       3,
		Name:     "Obstacle 3",
		Geometry: "Polygon((20 0, 30 0, 30 10, 20 10))",
	}

	tests := []struct {
		name          string
		body          string
		wantCode      int
		wantBody      string
		wantObstacles []Obstacle
	}{
		{
			name:          "valid obstacle",
			body:          `{"id":3,"name":"Obstacle 3","geometry":"Polygon((20 0, 30 0, 30 10, 20 10))"}`,
			wantCode:      http.StatusCreated,
			wantBody:      `{"id":3,"name":"Obstacle 3","geometry":"Polygon((20 0, 30 0, 30 10, 20 10))"}`,
			wantObstacles: append(fixtureObstacles(), newObstacle),
		},
		{
			name:          "invalid JSON",
			body:          `{"id":`,
			wantCode:      http.StatusBadRequest,
			wantObstacles: fixtureObstacles(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetObstacles(t)

			response := performRequest(router, http.MethodPost, "/obstacles", strings.NewReader(tt.body))

			assert.Equal(t, tt.wantCode, response.Code)
			assert.Equal(t, tt.wantObstacles, obstacles)
			if tt.wantBody != "" {
				assert.JSONEq(t, tt.wantBody, response.Body.String())
			}
		})
	}
}

func TestDeleteObstacle(t *testing.T) {
	router := setupRouter()
	tests := []struct {
		name          string
		path          string
		wantCode      int
		wantBody      string
		wantObstacles []Obstacle
	}{
		{
			name:          "existing obstacle",
			path:          "/obstacles/1",
			wantCode:      http.StatusNoContent,
			wantObstacles: fixtureObstacles()[1:],
		},
		{
			name:          "missing obstacle",
			path:          "/obstacles/3",
			wantCode:      http.StatusNotFound,
			wantBody:      `{"error":"obstacle not found"}`,
			wantObstacles: fixtureObstacles(),
		},
		{
			name:          "invalid ID",
			path:          "/obstacles/not-a-number",
			wantCode:      http.StatusBadRequest,
			wantBody:      `{"error":"invalid obstacle ID"}`,
			wantObstacles: fixtureObstacles(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetObstacles(t)

			response := performRequest(router, http.MethodDelete, tt.path, nil)

			assert.Equal(t, tt.wantCode, response.Code)
			assert.Equal(t, tt.wantObstacles, obstacles)
			if tt.wantBody != "" {
				assert.JSONEq(t, tt.wantBody, response.Body.String())
			}
		})
	}
}
