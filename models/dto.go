package models

type UserDetailsDto struct {
	User        *User         `json:"user"`
	Groups      []*Group      `json:"groups"`
	Roles       []*Role       `json:"roles"`
	Permissions []*Permission `json:"permissions"`
}

type TokenDto struct {
	Token string `json:"token"`
}
