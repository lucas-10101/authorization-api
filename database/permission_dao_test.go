package database

import (
	"context"
	"testing"

	"authorization-api/utils"
)

func TestPermissionDataAccessObject_FindAllByRoleId(t *testing.T) {
	utils.LoadEnv("../.env")

	dao := &PermissionDataAccessObject{
		Connection: GetDatabaseConnection(),
		Context:    context.Background(),
	}
	_, err := dao.FindAllByRoleId("dummy-role-id")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestPermissionDataAccessObject_FindAllByUserID(t *testing.T) {
	utils.LoadEnv("../.env")

	dao := &PermissionDataAccessObject{
		Connection: GetDatabaseConnection(),
		Context:    context.Background(),
	}
	_, err := dao.FindAllByUserID("dummy-user-id")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestPermissionDataAccessObject_FindAllByGroupId(t *testing.T) {
	utils.LoadEnv("../.env")
	dao := &PermissionDataAccessObject{
		Connection: GetDatabaseConnection(),
		Context:    context.Background(),
	}
	_, err := dao.FindAllByGroupId("dummy-group-id")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
