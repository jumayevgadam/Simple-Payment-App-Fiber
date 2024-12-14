package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	roleModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// Ensure RoleService implements the roles.Service interface.
var (
	_ roles.Service = (*RoleService)(nil)
)

// RoleService performs buisiness logic in role management.
type RoleService struct {
	repo database.DataStore
}

// NewRoleService creates and returns a new instance of RoleService.
func NewRoleService(repo database.DataStore) *RoleService {
	return &RoleService{repo: repo}
}

// AddRole service processes the request to create a new role and saves it to the database.
func (s *RoleService) AddRole(ctx context.Context, roleDTO roleModel.DTO) (int, error) {
	roleID, err := s.repo.RolesRepo().AddRole(ctx, roleDTO.ToStorage())
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return roleID, nil
}

// GetRole Service retrieves a role by its id from the database and returns it.
func (s *RoleService) GetRole(ctx context.Context, roleID int) (roleModel.DTO, error) {
	var roleDTO roleModel.DTO

	roleDAO, err := s.repo.RolesRepo().GetRole(ctx, roleID)
	if err != nil {
		return roleModel.DTO{}, errlst.ParseErrors(err)
	}
	roleDTO = roleDAO.ToServer()

	return roleDTO, nil
}

// GetRoles service fetches a list of all roles from database.
func (s *RoleService) GetRoles(ctx context.Context) ([]roleModel.DTO, error) {
	var roleDTOs []roleModel.DTO
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		roleDAOs, err := db.RolesRepo().GetRoles(ctx)
		if err != nil {
			return errlst.ErrNotFound
		}

		for _, role := range roleDAOs {
			roleDTOs = append(roleDTOs, role.ToServer())
		}

		return nil
	})

	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return roleDTOs, nil
}

// DeleteRole service processes deleting role action using its id and also deletes from database.
func (s *RoleService) DeleteRole(ctx context.Context, roleID int) error {
	// delete from database
	if err := s.repo.RolesRepo().DeleteRole(ctx, roleID); err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}

// UpdateRole service processes the new role data requests and saves it to the database.
func (s *RoleService) UpdateRole(ctx context.Context, roleDTO roleModel.DTO) (string, error) {
	var (
		resFromDB string
		err       error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		resFromDB, err = db.RolesRepo().UpdateRole(ctx, roleDTO.ToStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})
	if err != nil {
		return "", errlst.ParseErrors(err)
	}

	return resFromDB, nil
}
