package api

import (
	"reflect"
	"testing"
)

func TestAddTask(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		args    args
		want    Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTaskByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTaskByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTaskByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTasks(t *testing.T) {
	tests := []struct {
		name string
		want []*Task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTasks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveTaskById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveTaskById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RemoveTaskById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		args    args
		want    Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longUUID(); got != tt.want {
				t.Errorf("longUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_snowflakeUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snowflakeUUID(); got != tt.want {
				t.Errorf("snowflakeUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
