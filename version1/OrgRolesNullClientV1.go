package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type OrgRolesNullClientV1 struct {
}

func NewOrgRolesNullClientV1() *OrgRolesNullClientV1 {
	return &OrgRolesNullClientV1{}
}

func (c *OrgRolesNullClientV1) GetOrganizationUsers(ctx context.Context, correlationId string, orgId string) ([]string, error) {
	return make([]string, 0), nil
}

func (c *OrgRolesNullClientV1) GetOrganizationAdmins(ctx context.Context, correlationId string, orgId string) ([]string, error) {
	return make([]string, 0), nil
}

func (c *OrgRolesNullClientV1) GetOrganizationUserRoles(ctx context.Context, correlationId string, orgId string, paging *data.PagingParams) ([]*UserRolesV1, error) {
	return make([]*UserRolesV1, 0), nil
}

func (c *OrgRolesNullClientV1) GrantOrgRole(ctx context.Context, correlationId string, orgId string, userId string, role string) ([]string, error) {
	return []string{orgId + ":" + role}, nil
}

func (c *OrgRolesNullClientV1) RevokeOrgRole(ctx context.Context, correlationId string, orgId string, userId string, role string) ([]string, error) {
	return make([]string, 0), nil
}

func (c *OrgRolesNullClientV1) GrantDemoOrganizationUserRole(ctx context.Context, correlationId string, userId string, language string) (string, error) {
	return "demo", nil
}
