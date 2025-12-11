package httptransport

import (
	"encoding/json"
	"net/http"
	"time"

	"alemelomeza/silver-octo-parakeet/internal/domain/task"
	"alemelomeza/silver-octo-parakeet/internal/domain/user"
	taskusecase "alemelomeza/silver-octo-parakeet/internal/usecase/task"
)

type TaskHandler struct {
	CreateUC       *taskusecase.CreateTaskUseCase
	UpdateUC       *taskusecase.UpdateTaskUseCase
	DeleteUC       *taskusecase.DeleteTaskUseCase
	ListMyUC       *taskusecase.ListMyTasksUseCase
	ListAllUC      *taskusecase.ListAllTasksUseCase
	UpdateStatusUC *taskusecase.UpdateTaskStatusUseCase
	AddCommentUC   *taskusecase.AddCommentUseCase
}

func NewTaskHandler(
	create *taskusecase.CreateTaskUseCase,
	update *taskusecase.UpdateTaskUseCase,
	del *taskusecase.DeleteTaskUseCase,
	listMy *taskusecase.ListMyTasksUseCase,
	listAll *taskusecase.ListAllTasksUseCase,
	updateStatus *taskusecase.UpdateTaskStatusUseCase,
	addComment *taskusecase.AddCommentUseCase,
) *TaskHandler {
	return &TaskHandler{
		CreateUC:       create,
		UpdateUC:       update,
		DeleteUC:       del,
		ListMyUC:       listMy,
		ListAllUC:      listAll,
		UpdateStatusUC: updateStatus,
		AddCommentUC:   addComment,
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		AssignedTo  string `json:"assigned_to"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	due, _ := time.Parse(time.RFC3339, body.DueDate)

	role := r.Context().Value(CtxRole).(string)

	t, err := h.CreateUC.Execute(
		r.Context(),
		user.Role(role),
		body.Title,
		body.Description,
		due,
		body.AssignedTo,
	)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, t)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TaskID      string `json:"task_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		AssignedTo  string `json:"assigned_to"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	due, _ := time.Parse(time.RFC3339, body.DueDate)

	role := r.Context().Value(CtxRole).(string)

	t := &task.Task{
		ID:          body.TaskID,
		Title:       body.Title,
		Description: body.Description,
		DueDate:     due,
		AssignedTo:  body.AssignedTo,
	}

	err := h.UpdateUC.Execute(
		r.Context(),
		user.Role(role),
		t,
	)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TaskID string `json:"task_id"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	role := r.Context().Value(CtxRole).(string)

	err := h.DeleteUC.Execute(
		r.Context(),
		user.Role(role),
		body.TaskID,
	)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func (h *TaskHandler) ListMy(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(CtxUserID).(string)
	role := r.Context().Value(CtxRole).(string)

	tasks, err := h.ListMyUC.Execute(
		r.Context(),
		userID,
		user.Role(role),
	)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TaskID string `json:"task_id"`
		Status string `json:"status"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	userID := r.Context().Value(CtxUserID).(string)
	role := r.Context().Value(CtxRole).(string)

	err := h.UpdateStatusUC.Execute(
		r.Context(),
		user.Role(role),
		userID,
		body.TaskID,
		task.Status(body.Status),
	)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "updated"})
}

func (h *TaskHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TaskID  string `json:"task_id"`
		Comment string `json:"comment"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	userID := r.Context().Value(CtxUserID).(string)
	role := r.Context().Value(CtxRole).(string)

	err := h.AddCommentUC.Execute(
		r.Context(),
		user.Role(role),
		userID,
		body.TaskID,
		body.Comment,
	)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "comment added"})
}

func (h *TaskHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(CtxRole).(string)

	tasks, err := h.ListAllUC.Execute(
		r.Context(),
		user.Role(role),
	)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, tasks)
}


