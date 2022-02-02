package controller

import (
	"encoding/json"
	"fmt"
	"github.com/serverless-cloud-robot/robo9/internal/api"
	"net/http"
	"regexp"
)

type taskController struct {
	idPattern *regexp.Regexp
}

func (tc *taskController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/tasks" {
		switch r.Method {
		case http.MethodGet:
			tc.getAll(w, r)
		case http.MethodPost:
			tc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		match := tc.idPattern.FindStringSubmatch(r.URL.Path)
		if len(match) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id := match[1] // Need to convert this string to integer
		switch r.Method {
		case http.MethodGet:
			tc.get(id, w)
		case http.MethodPut:
			tc.put(id, w, r)
		case http.MethodDelete:
			tc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func NewTaskController() *taskController {
	return &taskController{
		idPattern: regexp.MustCompile(`^/tasks/(\d+)/?`),
	}
}
func (tc *taskController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(api.GetTasks(), w)
}

func (tc *taskController) get(id string, w http.ResponseWriter) {
	u, err := api.GetTaskByID(id) // TODO: Need to implement this function
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w) // TODO: Implement this function
}
func (tc *taskController) post(w http.ResponseWriter, r *http.Request) {
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

func (tc *taskController) put(id string, w http.ResponseWriter, r *http.Request) {
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

func (tc *taskController) delete(id string, w http.ResponseWriter) {
	err := api.RemoveTaskById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
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
