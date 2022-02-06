package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/serverless-cloud-robot/robo9/internal/api"
	"go.uber.org/zap"
	"net/http"
)

type taskController struct {
}

func (tc *taskController) GetAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(api.GetTasks(), w)
}

func (tc *taskController) Put(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	zap.S().Infof("ID %v", id)

	t, err := tc.parseRequest(r)
	if err != nil {
		zap.S().Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Task object"))
		return
	}
	if id != t.ID {
		zap.S().Info(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted Task must match ID of URL"))
		return
	}

	t, err = api.UpdateTask(t)
	if err != nil {
		zap.S().Info(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(t, w)
}

func (tc *taskController) Post(w http.ResponseWriter, r *http.Request) {
	u, err := tc.parseRequest(r)
	if err != nil {
		zap.S().Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Task object"))
		return
	}
	u, err = api.AddTask(u)
	if err != nil {
		zap.S().Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (tc *taskController) parseRequest(r *http.Request) (api.Task, error) {
	dec := json.NewDecoder(r.Body)
	var t api.Task
	err := dec.Decode(&t)
	if err != nil {
		zap.S().Info(err.Error())
		return api.Task{}, err
	}
	return t, nil
}

func (tc *taskController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	zap.S().Infof("ID %v", id)
	u, err := api.GetTaskByID(id)
	if err != nil {
		zap.S().Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)

}

func (tc *taskController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	zap.S().Infof("ID %v", id)
	err := api.RemoveTaskById(id)
	if err != nil {
		zap.S().Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewTaskController() *taskController {
	return &taskController{}
}
