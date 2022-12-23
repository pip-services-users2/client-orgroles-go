package version1

import "time"

type UserRolesV1 struct {
	Id         string    `json:"id"`
	Roles      []string  `json:"roles"`
	UpdateTime time.Time `json:"update_time"`
}

func NewUserRolesV1(id string, roles []string, updateTime time.Time) *UserRolesV1 {
	return &UserRolesV1{
		Id:         id,
		Roles:      roles,
		UpdateTime: updateTime,
	}
}
