package database

import (
	"authorization-api/models"
	"context"
	"database/sql"
)

type RoleDataAccessObject struct {
	Connection Connection
	Context    context.Context
}

func (dao *RoleDataAccessObject) FindAllByGroupId(id string) ([]*models.Role, error) {
	var roleList = []*models.Role{}
	var resultSet *sql.Rows
	var err error

	if resultSet, err = dao.Connection.QueryContext(
		dao.Context,
		`SELECT
			ROLES.ID,
			ROLES.NAME
		FROM
			GROUP_ROLES
			INNER JOIN ROLES ON ROLES.ID = GROUP_ROLES.ROLE_ID
		WHERE
			GROUP_ROLES.GROUP_ID = $1`,
		id,
	); err != nil {
		return nil, err
	}

	for resultSet.Next() {
		var role = &models.Role{}
		if err = resultSet.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}
		roleList = append(roleList, role)
	}

	return roleList, err
}
