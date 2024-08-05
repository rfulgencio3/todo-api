package handlers

import (
    "net/http"
    "strconv"
    "todo-api/models"

    "github.com/labstack/echo/v4"
)

var todos = []models.ToDo{}

// GetToDos godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ToDo
// @Router /todos [get]
func GetToDos(c echo.Context) error {
    return c.JSON(http.StatusOK, todos)
}

// CreateToDo godoc
// @Summary Create a new todo
// @Description Create a new todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body models.ToDo true "New ToDo"
// @Success 201 {object} models.ToDo
// @Router /todos [post]
func CreateToDo(c echo.Context) error {
    todo := new(models.ToDo)
    if err := c.Bind(todo); err != nil {
        return err
    }
    todo.ID = len(todos) + 1
    todos = append(todos, *todo)
    return c.JSON(http.StatusCreated, todo)
}

// UpdateToDo godoc
// @Summary Update an existing todo
// @Description Update an existing todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "ToDo ID"
// @Param todo body models.ToDo true "Updated ToDo"
// @Success 200 {object} models.ToDo
// @Router /todos/{id} [put]
func UpdateToDo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    todo := new(models.ToDo)
    if err := c.Bind(todo); err != nil {
        return err
    }
    for i, t := range todos {
        if t.ID == id {
            todos[i] = *todo
            return c.JSON(http.StatusOK, todo)
        }
    }
    return c.JSON(http.StatusNotFound, nil)
}

// DeleteToDo godoc
// @Summary Delete a todo
// @Description Delete a todo
// @Tags todos
// @Param id path int true "ToDo ID"
// @Success 204
// @Router /todos/{id} [delete]
func DeleteToDo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, t := range todos {
        if t.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            return c.NoContent(http.StatusNoContent)
        }
    }
    return c.JSON(http.StatusNotFound, nil)
}
