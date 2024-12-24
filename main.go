// main.go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var tasks []Task

// Create a task
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task.ID = fmt.Sprintf("%d", len(tasks)+1) // Simple ID assignment
	tasks = append(tasks, task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// Get all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Get a single task by ID
func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, task := range tasks {
		if task.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// Update a task by ID
func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedTask Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, task := range tasks {
		if task.ID == params["id"] {
			tasks[index].Title = updatedTask.Title
			tasks[index].Description = updatedTask.Description
			tasks[index].Completed = updatedTask.Completed
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks[index])
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

// Delete a task by ID
func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, task := range tasks {
		if task.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// Main function to setup routes and start the server
func main() {
	r := mux.NewRouter()

	// Seed some initial tasks for demo purposes
	tasks = append(tasks, Task{ID: "1", Title: "Go to the gym", Description: "Workout session", Completed: false})
	tasks = append(tasks, Task{ID: "2", Title: "Buy groceries", Description: "Get fruits and vegetables", Completed: false})

	// Define routes
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	// Start the server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
