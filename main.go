//go:generate goagen bootstrap -d github.com/jarifibrahim/todo/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jarifibrahim/todo/app"
)

func main() {
	// Create service
	service := goa.New("Todo List")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "todo" controller
	c := NewTodoController(service)
	app.MountTodoController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
