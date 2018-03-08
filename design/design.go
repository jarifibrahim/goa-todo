package design

import (
	d "github.com/goadesign/goa/design"
	a "github.com/goadesign/goa/design/apidsl"
)

var _ = a.API("Todo List", func() {
	a.Title("A Todo List api")
	a.Description("A simple goa service")
	a.Scheme("http")
	a.Host("localhost:8080")
	a.BasePath("/api/v1")
})

var _ = a.Resource("todo", func() {
	a.BasePath("/todos")

	a.Action("create", func() {
		a.Description("Create a Todo item")
		a.Routing(a.POST(""))
		a.Payload(todoPayload)
		a.Response(d.Created, todoMedia)
		a.Response(d.BadRequest)
	})

	a.Action("delete", func() {
		a.Description("Delete a Todo item")
		a.Routing(a.DELETE("/:id"))
		a.Params(func() {
			a.Param("id", d.Integer, "ID of item to be deleted")
		})
		a.Response(d.OK)
		a.Response(d.NotFound)
	})

	a.Action("list", func() {
		a.Description("List all Todos")
		a.Routing(a.GET(""))
		a.Response(d.OK, a.ArrayOf(todoMedia))
		a.Response(d.NotFound)
	})

	a.Action("update", func() {
		a.Description("Update a Todo item")
		a.Routing(a.PATCH("/:id"))
		a.Params(func() {
			a.Param("id", d.Integer, "ID of the todo item")
		})
		a.Payload(todoPayload)
		a.Response(d.OK, todoMedia)
		a.Response(d.NotFound)
		a.Response(d.BadRequest)
	})

	a.Action("show", func() {
		a.Description("Show a ToDo item")
		a.Routing(a.GET("/:id"))
		a.Params(func() {
			a.Param("id", d.Integer, "id")
		})
		a.Response(d.OK, todoMedia)
		a.Response(d.NotFound)
		a.Response(d.BadRequest)
	})
})

var todoPayload = a.Type("todoPayload", func() {
	a.Attribute("title", d.String, "Title of ToDO")
	a.Attribute("description", d.String, "Description of ToDo")
	a.Required("title", "description")
})

var todoMedia = a.MediaType("application/vnd.todo+json", func() {
	a.Reference(todoPayload)
	a.TypeName("todo")
	a.Attributes(func() {
		a.Attribute("ID", d.Integer, "Unique todo ID")
		a.Attribute("title")
		a.Attribute("description")
		a.Required("ID", "title", "description")
	})

	a.View("default", func() {
		a.Attribute("ID")
		a.Attribute("title")
		a.Attribute("description")
	})
})
