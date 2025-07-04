package models

import "crypto/rsa"

type Tenant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	TenantId string `json:"tenant_id"`
}

type Group struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TenantId string `json:"tenant_id"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserGroup struct {
	UserID  string `json:"user_id"`
	GroupID string `json:"group_id"`
}

type GroupRole struct {
	GroupID string `json:"group_id"`
	RoleID  string `json:"role_id"`
}

type Permission struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Resource    string `json:"resource"`
	Scope       string `json:"scope"`
}

func (p *Permission) ToScopedResource() string {
	return p.Resource + ":" + p.Scope
}

type RolePermission struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

type SigningKeyGroup string

var (
	JWT_KEY_GROUP SigningKeyGroup = SigningKeyGroup("JWT")
)

type SigningKey struct {
	Kid           string          `json:"-"`
	CreatedAt     string          `json:"-"`
	PrivateKey    string          `json:"-"`
	RsaPrivateKey *rsa.PrivateKey `json:"-" db:"-"`
	Active        bool            `json:"-"`
	KeyGroup      SigningKeyGroup `json:"-"`
}
