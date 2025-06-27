package database

import (
	"context"
	"testing"

	"authorization-api/utils"
)

func TestUserDataAccessObject_FindUserByUsernameAndTenantId(t *testing.T) {
	utils.LoadEnv("../.env")
	dao := UserDataAccessObject{
		Connection: GetDatabaseConnection(),
		Context:    context.Background(),
	}
	_, err := dao.FindUserByUsernameAndTenantId("dummy-username", "dummy-password", "dummy-tenant-id")
	if err == nil {
		t.Errorf("expected an error, got nil")
	}

	if err.Error() != "user not found" {
		t.Errorf("expected 'user not found' error, got %v", err)
	}
}
