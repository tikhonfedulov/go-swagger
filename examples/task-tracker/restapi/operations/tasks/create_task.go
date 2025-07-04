// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateTaskHandlerFunc turns a function with the right signature into a create task handler
type CreateTaskHandlerFunc func(CreateTaskParams, any) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTaskHandlerFunc) Handle(params CreateTaskParams, principal any) middleware.Responder {
	return fn(params, principal)
}

// CreateTaskHandler interface for that can handle valid create task params
type CreateTaskHandler interface {
	Handle(CreateTaskParams, any) middleware.Responder
}

// NewCreateTask creates a new http.Handler for the create task operation
func NewCreateTask(ctx *middleware.Context, handler CreateTaskHandler) *CreateTask {
	return &CreateTask{Context: ctx, Handler: handler}
}

/*
	CreateTask swagger:route POST /tasks tasks createTask

Creates a 'Task' object.

Allows for creating a task.
This operation requires authentication so that we know which user
created the task.
*/
type CreateTask struct {
	Context *middleware.Context
	Handler CreateTaskHandler
}

func (o *CreateTask) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateTaskParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal any
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
