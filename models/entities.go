package models

type Tenant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
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

type SigningKey struct {
	Kid        string `json:"-"`
	CreatedAt  string `json:"-"`
	PrivateKey string `json:"-"`
}
