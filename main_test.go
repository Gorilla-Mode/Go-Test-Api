package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var testObstacles = []Obstacle{
	{
		ID:       1,
		Name:     "Obstacle 1",
		Geometry: "Polygon((0 0, 10 0, 10 10, 0 10))",
	},
	{
		ID:       2,
		Name:     "Obstacle 2",
		Geometry: "Polygon((10 0, 20 0, 20 10, 10 10))",
	},
}

func newTestRouter() http.Handler {
	gin.SetMode(gin.TestMode)

	// Copy the slice so each test starts with fresh data.
	obstacles = append([]Obstacle(nil), testObstacles...)

	return setupRouter()
}

func performRequest(
	router http.Handler,
	method string,
	path string,
	body string,
) *httptest.ResponseRecorder {
	request := httptest.NewRequest(
		method,
		path,
		strings.NewReader(body),
	)

	if body != "" {
		request.Header.Set("Content-Type", "application/json")
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	return response
}

func TestGetObstacles(t *testing.T) {
	router := newTestRouter()

	response := performRequest(
		router,
		http.MethodGet,
		"/obstacles",
		"",
	)

	require.Equal(t, http.StatusOK, response.Code)
	require.JSONEq(t, `[
		{
			"id": 1,
			"name": "Obstacle 1",
			"geometry": "Polygon((0 0, 10 0, 10 10, 0 10))"
		},
		{
			"id": 2,
			"name": "Obstacle 2",
			"geometry": "Polygon((10 0, 20 0, 20 10, 10 10))"
		}
	]`, response.Body.String())
}

func TestGetObstacle(t *testing.T) {
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
			wantBody: `{
				"id": 1,
				"name": "Obstacle 1",
				"geometry": "Polygon((0 0, 10 0, 10 10, 0 10))"
			}`,
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := newTestRouter()

			response := performRequest(
				router,
				http.MethodGet,
				test.path,
				"",
			)

			require.Equal(t, test.wantCode, response.Code)
			require.JSONEq(t, test.wantBody, response.Body.String())
		})
	}
}

func TestPostObstacle(t *testing.T) {
	router := newTestRouter()

	body := `{
		"id": 3,
		"name": "Obstacle 3",
		"geometry": "Polygon((20 0, 30 0, 30 10, 20 10))"
	}`

	response := performRequest(
		router,
		http.MethodPost,
		"/obstacles",
		body,
	)

	require.Equal(t, http.StatusCreated, response.Code)
	require.JSONEq(t, body, response.Body.String())
	require.Len(t, obstacles, 3)
	require.Equal(t, uint64(3), obstacles[2].ID)
}

func TestPostObstacleInvalidJSON(t *testing.T) {
	router := newTestRouter()

	response := performRequest(
		router,
		http.MethodPost,
		"/obstacles",
		`{"id":`,
	)

	require.Equal(t, http.StatusBadRequest, response.Code)
	require.Len(t, obstacles, 2)
}

func TestDeleteObstacle(t *testing.T) {
	router := newTestRouter()

	response := performRequest(
		router,
		http.MethodDelete,
		"/obstacles/1",
		"",
	)

	require.Equal(t, http.StatusNoContent, response.Code)
	require.Empty(t, response.Body.String())
	require.Equal(t, testObstacles[1:], obstacles)
}
