package main

import (
	"github.com/goadesign/goa"
	"github.com/jarifibrahim/todo/app"
)

// TodoController implements the todo resource.
type TodoController struct {
	*goa.Controller
}

// NewTodoController creates a todo controller.
func NewTodoController(service *goa.Service) *TodoController {
	return &TodoController{Controller: service.NewController("TodoController")}
}

// Create runs the create action.
func (c *TodoController) Create(ctx *app.CreateTodoContext) error {
	// TodoController_Create: start_implement

	// Put your logic here

	return nil
	// TodoController_Create: end_implement
}

// Delete runs the delete action.
func (c *TodoController) Delete(ctx *app.DeleteTodoContext) error {
	// TodoController_Delete: start_implement

	// Put your logic here

	return nil
	// TodoController_Delete: end_implement
}

// List runs the list action.
func (c *TodoController) List(ctx *app.ListTodoContext) error {
	// TodoController_List: start_implement

	// Put your logic here
	return nil
	// TodoController_List: end_implement
}

// Show runs the show action.
func (c *TodoController) Show(ctx *app.ShowTodoContext) error {
	// TodoController_Show: start_implement

	// Put your logic here

	res := &app.Todo{}
	return ctx.OK(res)
	// TodoController_Show: end_implement
}

// Update runs the update action.
func (c *TodoController) Update(ctx *app.UpdateTodoContext) error {
	// TodoController_Update: start_implement

	// Put your logic here

	res := &app.Todo{}
	return ctx.OK(res)
	// TodoController_Update: end_implement
}
