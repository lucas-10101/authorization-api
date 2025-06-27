package database

import (
	"context"
	"testing"

	"authorization-api/utils"
)

func TestRoleDataAccessObject_FindAllByGroupId(t *testing.T) {
	utils.LoadEnv("../.env")
	dao := &RoleDataAccessObject{
		Connection: GetDatabaseConnection(),
		Context:    context.Background(),
	}
	_, err := dao.FindAllByGroupId("dummy-group-id")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
