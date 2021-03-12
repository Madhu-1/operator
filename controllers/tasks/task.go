/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tasks

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TaskError holds the known error and the Task Name
type TaskError struct {
	ErrorCode codes.Code
	Name      string
}

// TaskSpec is the specification for each Task
type TaskSpec struct {
	Name        string
	Task        Task
	KnownErrors []TaskError
}

// Task is a specific task to be done by controller
type Task interface {
	Run() error
}

// RunAll executes all the Task in the given list of TaskSpec
func RunAll(tasks []*TaskSpec) (string, error) {
	for i, task := range tasks {
		if err := task.Task.Run(); err != nil {
			foundError := false
			sc, ok := status.FromError(err)
			if !ok {
				// This is not gRPC error. The operation must have failed before gRPC
				// method was called, otherwise we would get gRPC error.
				return task.Name, err
			}
			// check for next task error message if next task can be continued
			// for the current error continue it
			if i < len(tasks) {
				for _, e := range tasks[i+1].KnownErrors {
					if task.Name == e.Name && sc.Code() == e.ErrorCode {
						foundError = true
						break
					}
				}
			}
			if !foundError {
				return task.Name, err
			}
		}
	}
	return "", nil
}
