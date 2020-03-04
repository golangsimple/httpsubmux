package httpsubmux


import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServeMux(t *testing.T) {
	tt := []struct {
		method string
		target string
		body io.Reader
		responseCode int
	}{
		{
			method: http.MethodGet,
			target: "/",
			responseCode: http.StatusNotFound,
		},
		{
			method: http.MethodGet,
			target: "/api/",
			responseCode: http.StatusNotFound,
		},
		{
			method: http.MethodGet,
			target: "/api/tasks/",
			responseCode: http.StatusNotFound,
		},
		{
			method: http.MethodGet,
			target: "/api/tasks/ok",
			responseCode: http.StatusOK,
		},
		{
			method: http.MethodGet,
			target: "/api/tasks/notes/ok",
			responseCode: http.StatusOK,
		},
	}

	handler := SetupMux()

	for _, test := range tt {
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, httptest.NewRequest(test.method, test.target, nil))
		if response.Code != test.responseCode {
			t.Errorf("Incorrect response code %v", response.Code)
		}
	}
}

func SetupMux() ServeMux {
	handler := http.NewServeMux()

	apiHandler := NewServeMux("", "/api/")
	handler.Handle(apiHandler.Route, apiHandler)

	tasksHandler := NewServeMux(apiHandler.Route, "tasks/")
	apiHandler.Handle(tasksHandler.Route, tasksHandler)

	tasksHandler.HandleFunc(tasksHandler.Route + "ok", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	notesHandler := NewServeMux(tasksHandler.Route, "notes/")
	tasksHandler.Handle(notesHandler.Route, notesHandler)

	notesHandler.HandleFunc(notesHandler.Route + "ok", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	return handler
}
