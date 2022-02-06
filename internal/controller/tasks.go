package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/serverless-cloud-robot/robo9/internal/api"
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
	fmt.Println(id)

	t, err := tc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Task object"))
		fmt.Println(err.Error())
		return
	}
	if id != t.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted Task must match ID of URL"))
		fmt.Println(err.Error())
		return
	}

	t, err = api.UpdateTask(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	encodeResponseAsJSON(t, w)
}

func (tc *taskController) Post(w http.ResponseWriter, r *http.Request) {
	u, err := tc.parseRequest(r) // TODO: Implement this function
	fmt.Println("Parsed :")
	fmt.Println(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Task object"))
		return
	}
	u, err = api.AddTask(u)
	if err != nil {
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
	fmt.Println(t)
	if err != nil {
		return api.Task{}, err
	}
	return t, nil
}

func (tc *taskController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	u, err := api.GetTaskByID(id) // TODO: Need to implement this function
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)

}

func (tc *taskController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := api.RemoveTaskById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewTaskController() *taskController {
	return &taskController{}
}
