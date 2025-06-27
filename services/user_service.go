package services

import (
	"authorization-api/database"
	"authorization-api/models"
	"context"
)

type UserService struct {
	Connection database.Connection
	Context    context.Context
}

func (s *UserService) GetUserDetails(userId string) (*models.UserDetailsDto, error) {

	userDao := &database.UserDataAccessObject{
		Connection: s.Connection,
		Context:    s.Context,
	}

	groupDao := &database.GroupDataAccessObject{
		Connection: s.Connection,
		Context:    s.Context,
	}

	roleDao := &database.RoleDataAccessObject{
		Connection: s.Connection,
		Context:    s.Context,
	}

	permissionDao := &database.PermissionDataAccessObject{
		Connection: s.Connection,
		Context:    s.Context,
	}

	var err error
	var groups []*models.Group
	var roles []*models.Role
	var permissions []*models.Permission
	var user *models.User

	user, err = userDao.FindUserById(userId)
	if err != nil {
		return nil, err
	}

	groups, err = groupDao.FindAllByUserId(userId)
	if err != nil {
		return nil, err
	}

	roles = make([]*models.Role, 0)
	for _, group := range groups {
		groupRoles, err := roleDao.FindAllByGroupId(group.ID)
		if err != nil {
			return nil, err
		}
		roles = append(roles, groupRoles...)
	}

	permissions, err = permissionDao.FindAllByUserID(userId)
	if err != nil {
		return nil, err
	}

	userDto := &models.UserDetailsDto{
		User:        user,
		Groups:      groups,
		Roles:       roles,
		Permissions: permissions,
	}

	return userDto, err
}
