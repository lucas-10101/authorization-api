package database

import (
	"authorization-api/models"
	"context"
	"database/sql"
)

type GroupDataAccessObject struct {
	Connection Connection
	Context    context.Context
}

func (dao *GroupDataAccessObject) FindAllByUserId(id string) ([]*models.Group, error) {

	var groupList = []*models.Group{}
	var resultSet *sql.Rows
	var err error

	if resultSet, err = dao.Connection.QueryContext(
		dao.Context,
		`SELECT
			ID,
			NAME,
			TENANT_ID
		FROM
			USER_GROUPS
			INNER JOIN GROUPS ON USER_GROUPS.GROUP_ID = GROUPS.ID
		WHERE
			USER_GROUPS.USER_ID = $1`,
		id,
	); err != nil {
		return nil, err
	}

	for resultSet.Next() {
		var group = &models.Group{}
		if err = resultSet.Scan(&group.ID, &group.Name, &group.TenantId); err != nil {
			return nil, err
		}
		groupList = append(groupList, group)
	}

	return groupList, err
}
