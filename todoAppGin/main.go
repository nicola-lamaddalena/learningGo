package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Task{}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", postTodo)
	router.PUT("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodoByID)

	router.Run("localhost:8000")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, todo := range todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusCreated, todo)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func postTodo(c *gin.Context) {
	var newTodo Task

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var newTodo Task

	for index, todo := range todos {
		if todo.ID == id {

			if err := c.BindJSON(&newTodo); err != nil {
				return
			}
			// remove the todo from the slice
			todos = append(todos[:index], todos[index+1:]...)
			// insert the updated version
			todos = append(todos, newTodo)
			c.IndentedJSON(http.StatusAccepted, newTodo)
		}
	}
}

func deleteTodoByID(c *gin.Context) {
	id := c.Param("id")

	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			c.IndentedJSON(http.StatusAccepted, todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
