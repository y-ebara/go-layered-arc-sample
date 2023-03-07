package rest

import (
	"encoding/json"
	"net/http"
	"test_app2/models"
	"test_app2/service"
	"time"

	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	sevice service.TodoService
}

func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{sevice: service}
}

func (h *TodoHandler) Routes(mux chi.Router) chi.Router {
	mux.Post("/todo/tasks", h.CreateTask)
	mux.Put("/todo/task/{id}", h.UpdateTask)
	return mux
}

type CreateTaskParams struct{
	description string
	priority int
	start time.Time
	limit time.Time
}

func (h *TodoHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	param := &CreateTaskParams{}
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	task, err := h.sevice.CreateNewTask(r.Context(), param.description, models.Priority(param.priority), models.TaskPeriod{
		Start: param.start,
		Limit: param.limit,
	})
	if err != nil {
		w.Write([]byte("error"))
		return
	}

	resp := map[string]interface{}{
		"task": task,
		"ok": true,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		 w.Write([]byte("error"))
		 return
	}
	w.Write(json)
}

func (h *TodoHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	// TODO: impl if needed
}
