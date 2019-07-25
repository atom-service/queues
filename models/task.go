package models

import "github.com/grpcbrick/queues/standard"

// Task 任务
type Task struct {
	ID            uint64 `json:"ID"`
	Next          uint64 `json:"Next"`
	Prior         uint64 `json:"Prior"`
	Owner         uint64 `json:"Owner"`
	State         string `json:"State"`
	Input         string `json:"Input"`
	Output        string `json:"Output"`
	Channel       string `son:"Channel"`
	HashCode      string `son:"HashCode"`
	RetryCount    uint64 `json:"RetryCount"`
	CreateTime    string `json:"CreateTime"`
	UpdateTime    string `json:"UpdateTime"`
	RetryMaxLimit uint64 `json:"RetryMaxLimit"`
}

// LoadProtoStruct LoadProtoStruct
func (srv *Task) LoadProtoStruct(task *standard.Task) {
	srv.ID = task.ID
	srv.Next = task.Next
	srv.Prior = task.Prior
	srv.Owner = task.Owner
	srv.State = task.State
	srv.Input = task.Input
	srv.Output = task.Output
	srv.Channel = task.Channel
	srv.HashCode = task.HashCode
	srv.RetryCount = task.RetryCount
	srv.CreateTime = task.CreateTime
	srv.UpdateTime = task.UpdateTime
	srv.RetryMaxLimit = task.RetryMaxLimit
}

// OutProtoStruct OutProtoStruct
func (srv *Task) OutProtoStruct() *standard.Task {
	task := new(standard.Task)
	task.ID = srv.ID
	task.Next = srv.Next
	task.Prior = srv.Prior
	task.Owner = srv.Owner
	task.State = srv.State
	task.Input = srv.Input
	task.Output = srv.Output
	task.Channel = srv.Channel
	task.HashCode = srv.HashCode
	task.RetryCount = srv.RetryCount
	task.CreateTime = srv.CreateTime
	task.UpdateTime = srv.UpdateTime
	task.RetryMaxLimit = srv.RetryMaxLimit
	return task
}
