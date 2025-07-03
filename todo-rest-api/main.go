package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}

var (
	tasks []Task
	users []User
)

func main() {

	users = []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
	tasks = []Task{
		{ID: 1, Title: "Task 1", Description: "Do something", UserId: 1},
		{ID: 2, Title: "Task 2", Description: "Do more", UserId: 1},
		{ID: 3, Title: "Task 3", Description: "Bob's task", UserId: 2},
	}

	router := mux.NewRouter()

	// Define the endpoints for CRUD operations
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/user/{user_id}", getTasksByUser).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.NotFound(w, r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = len(users) + 1
	users = append(users, user)

	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	http.NotFound(w, r)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func getTasksByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.Atoi(params["user_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userTasks []Task

	for _, task := range tasks {
		if task.UserId == userId {
			userTasks = append(userTasks, task)
		}
	}
	if len(userTasks) > 0 {
		json.NewEncoder(w).Encode(userTasks)
		return
	}
	http.NotFound(w, r)
}

// func getTasksByUser(w http.ResponseWriter, r *http.Request) {
//     params := mux.Vars(r)
//     userId, err := strconv.Atoi(params["user_id"])
//     if err != nil {
//         http.Error(w, "Invalid user_id", http.StatusBadRequest)
//         return
//     }

//     // Проверяем, существует ли пользователь
//     userExists := false
//     for _, user := range users {
//         if user.ID == userId {
//             userExists = true
//             break
//         }
//     }
//     if !userExists {
//         http.NotFound(w, r)  // Пользователь не найден
//         return
//     }

//     // Ищем задачи пользователя
//     var userTasks []Task
//     for _, task := range tasks {
//         if task.UserId == userId {
//             userTasks = append(userTasks, task)
//         }
//     }

//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(userTasks)  // Возвращаем [] или список задач
// }

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.NotFound(w, r)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = len(tasks) + 1
	tasks = append(tasks, task)

	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedTask Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	http.NotFound(w, r)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}
