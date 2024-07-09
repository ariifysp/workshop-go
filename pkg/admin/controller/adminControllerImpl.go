package controller

import (
	_adminService "github/ariifysp/workshop-go/pkg/admin/service"
)

type adminControllerImpl struct {
	adminService _adminService.AdminService
}

func NewAdminControllerImpl(
	adminService _adminService.AdminService,
) AdminController {
	return &adminControllerImpl{adminService}
}
