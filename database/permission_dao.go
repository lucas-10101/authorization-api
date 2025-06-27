package database

import (
	"authorization-api/models"
	"context"
	"database/sql"
)

type PermissionDataAccessObject struct {
	Connection Connection
	Context    context.Context
}

func (dao *PermissionDataAccessObject) FindAllByRoleId(roleId string) ([]*models.Permission, error) {

	var permissionList = []*models.Permission{}
	var resultSet *sql.Rows
	var err error
	if resultSet, err = dao.Connection.QueryContext(
		dao.Context,
		`SELECT
			PERMISSIONS.ID,
			PERMISSIONS.NAME
		FROM
			ROLE_PERMISSIONS
			INNER JOIN PERMISSIONS ON PERMISSIONS.ID = ROLE_PERMISSIONS.PERMISSION_ID
		WHERE
			ROLE_PERMISSIONS.ROLE_ID = $1`,
		roleId,
	); err != nil {
		return nil, err
	}

	for resultSet.Next() {
		var permission = &models.Permission{}
		if err = resultSet.Scan(&permission.ID, &permission.Name); err != nil {
			return nil, err
		}
		permissionList = append(permissionList, permission)
	}

	return permissionList, nil
}

func (dao *PermissionDataAccessObject) FindAllByUserID(userId string) ([]*models.Permission, error) {

	var permissionList = []*models.Permission{}
	var resultSet *sql.Rows
	var err error
	if resultSet, err = dao.Connection.QueryContext(
		dao.Context,
		`SELECT
			PERMISSIONS.ID,
			PERMISSIONS.NAME
		FROM
			USER_GROUPS
			INNER JOIN GROUPS ON USER_GROUPS.GROUP_ID = GROUPS.ID
			INNER JOIN GROUP_ROLES ON GROUP_ROLES.GROUP_ID = USER_GROUPS.GROUP_ID
			INNER JOIN ROLES ON ROLES.ID = GROUP_ROLES.ROLE_ID
			INNER JOIN ROLE_PERMISSIONS ON ROLE_PERMISSIONS.ROLE_ID = ROLES.ID
			INNER JOIN PERMISSIONS ON PERMISSIONS.ID = ROLE_PERMISSIONS.PERMISSION_ID
		WHERE
			USER_GROUPS.USER_ID = $1`,
		userId,
	); err != nil {
		return nil, err
	}

	for resultSet.Next() {
		var permission = &models.Permission{}
		if err = resultSet.Scan(&permission.ID, &permission.Name); err != nil {
			return nil, err
		}
		permissionList = append(permissionList, permission)
	}

	return permissionList, err
}

func (dao *PermissionDataAccessObject) FindAllByGroupId(groupId string) ([]*models.Permission, error) {

	var permissionList = []*models.Permission{}
	var resultSet *sql.Rows
	var err error
	if resultSet, err = dao.Connection.QueryContext(
		dao.Context,
		`SELECT
			PERMISSIONS.ID,
			PERMISSIONS.NAME
		FROM
			GROUP_ROLES
			INNER JOIN ROLES ON ROLES.ID = GROUP_ROLES.ROLE_ID
			INNER JOIN ROLE_PERMISSIONS ON ROLE_PERMISSIONS.ROLE_ID = ROLES.ID
			INNER JOIN PERMISSIONS ON PERMISSIONS.ID = ROLE_PERMISSIONS.PERMISSION_ID
		WHERE
			GROUP_ROLES.GROUP_ID = $1`,
		groupId,
	); err != nil {
		return nil, err
	}

	for resultSet.Next() {
		var permission = &models.Permission{}
		if err = resultSet.Scan(&permission.ID, &permission.Name); err != nil {
			return nil, err
		}
		permissionList = append(permissionList, permission)
	}

	return permissionList, err
}
