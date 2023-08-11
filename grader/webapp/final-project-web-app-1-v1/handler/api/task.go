package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TaskAPI interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	CreateNewTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTaskCategory(w http.ResponseWriter, r *http.Request)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskService service.TaskService) *taskAPI {
	return &taskAPI{taskService}
}

func (t *taskAPI) GetTask(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	userId := r.Context().Value("id").(string)
	intUserId, _ := strconv.Atoi(userId)

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	taskId := r.URL.Query().Get("task_id")
	intTaskId, _ := strconv.Atoi(taskId)
	if taskId == "" {
		data, err := t.taskService.GetTasks(r.Context(), intUserId)
		if err != nil {
			if err != nil {
				w.WriteHeader(500)
				json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
				return
			}
		}
		w.WriteHeader(200)
		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)
	} else {
		data, err := t.taskService.GetTaskByID(r.Context(), intTaskId)
		if err != nil {
			if err != nil {
				w.WriteHeader(500)
				json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
				return
			}
		}
		w.WriteHeader(200)
		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)
	}

}

func (t *taskAPI) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}

	// TODO: answer here
	if task.Title == "" || task.Description == "" || task.CategoryID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}

	userId := r.Context().Value("id").(string)
	intUserId, _ := strconv.Atoi(userId)

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var tsk = entity.Task{
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
		UserID:      intUserId,
	}
	data, err := t.taskService.StoreTask(r.Context(), &tsk)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(201)
	jsonData, _ := json.Marshal(map[string]interface{}{
		"user_id": intUserId,
		"task_id": data.ID,
		"message": "success create new task",
	})
	w.Write(jsonData)
}

func (t *taskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	userId := r.Context().Value("id").(string)
	intUserId, _ := strconv.Atoi(userId)
	taskId := r.URL.Query().Get("task_id")
	intTaskId, _ := strconv.Atoi(taskId)

	err := t.taskService.DeleteTask(r.Context(), intTaskId)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(200)
	jsonData, _ := json.Marshal(map[string]interface{}{
		"user_id": intUserId,
		"task_id": intTaskId,
		"message": "success delete task",
	})
	w.Write(jsonData)
}

func (t *taskAPI) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	// TODO: answer here
	userId := r.Context().Value("id").(string)
	intUserId, _ := strconv.Atoi(userId)

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var tsk = entity.Task{
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
		UserID:      intUserId,
		ID:          task.ID,
	}
	data, err := t.taskService.UpdateTask(r.Context(), &tsk)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(200)
	jsonData, _ := json.Marshal(map[string]interface{}{
		"user_id": intUserId,
		"task_id": data.ID,
		"message": "success update task",
	})
	w.Write(jsonData)
}

func (t *taskAPI) UpdateTaskCategory(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var updateTask = entity.Task{
		ID:         task.ID,
		CategoryID: task.CategoryID,
		UserID:     int(idLogin),
	}

	_, err = t.taskService.UpdateTask(r.Context(), &updateTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userId,
		"task_id": task.ID,
		"message": "success update task category",
	})
}
