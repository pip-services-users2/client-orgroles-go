package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IOrgRolesClientV1 interface {
	GetOrganizationUsers(ctx context.Context, correlationId string, orgId string) ([]string, error)

	GetOrganizationAdmins(ctx context.Context, correlationId string, orgId string) ([]string, error)

	GetOrganizationUserRoles(ctx context.Context, correlationId string, orgId string, paging *data.PagingParams) ([]*UserRolesV1, error)

	GrantOrgRole(ctx context.Context, correlationId string, orgId string, userId string, role string) ([]string, error)

	RevokeOrgRole(ctx context.Context, correlationId string, orgId string, userId string, role string) ([]string, error)

	GrantDemoOrganizationUserRole(ctx context.Context, correlationId string, userId string, language string) (string, error)
}
