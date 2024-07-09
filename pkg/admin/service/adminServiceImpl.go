package service

import (
	_adminRepository "github/ariifysp/workshop-go/pkg/admin/repository"
)

type adminServiceImpl struct {
	adminRepository _adminRepository.AdminRepository
}

func NewAdminServiceImpl(
	adminRepository _adminRepository.AdminRepository,
) AdminService {
	return &adminServiceImpl{adminRepository}
}
