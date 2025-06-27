package database

import (
	"context"
	"testing"

	"authorization-api/utils"
)

func TestGroupDataAccessObject_FindAllByUserId(t *testing.T) {
	utils.LoadEnv("../.env")
	dao := &GroupDataAccessObject{
		Connection: GetDatabaseConnection(),
		Context:    context.Background(),
	}
	_, err := dao.FindAllByUserId("dummy-user-id")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
