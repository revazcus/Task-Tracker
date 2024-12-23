package router

import (
	"net/http"
	roleRest "task-tracker/adapters/controllers/rest/role"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type RoleRouter struct {
	controller *roleRest.RoleController
}

func NewRoleRouter(controller *roleRest.RoleController) *RoleRouter {
	return &RoleRouter{
		controller: controller,
	}
}

func (r *RoleRouter) RegisterRoutes(server restServerInterface.Server) {
	server.RegisterPublicRoute(http.MethodGet, "v1/spec", r.controller.GetRoleById)
	server.RegisterPublicRoute(http.MethodPost, "v1/spec/create", r.controller.CreateRole)
	server.RegisterPublicRoute(http.MethodPut, "v1/spec/update", r.controller.UpdateRole)
	server.RegisterPublicRoute(http.MethodDelete, "v1/spec", r.controller.DeleteRoleById)
}
