package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type OrgRolesCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewOrgRolesCommandableGrpcClientV1() *OrgRolesCommandableGrpcClientV1 {
	return NewOrgRolesCommandableGrpcClientV1ithConfig(nil)
}

func NewOrgRolesCommandableGrpcClientV1ithConfig(config *cconf.ConfigParams) *OrgRolesCommandableGrpcClientV1 {
	c := &OrgRolesCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/org_roles"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *OrgRolesCommandableGrpcClientV1) GetOrganizationUsers(ctx context.Context, correlationId string, orgId string) ([]string, error) {
	res, err := c.CallCommand(ctx, "get_organization_users", correlationId, data.NewAnyValueMapFromTuples(
		"org_id", orgId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *OrgRolesCommandableGrpcClientV1) GetOrganizationAdmins(ctx context.Context, correlationId string, orgId string) ([]string, error) {
	res, err := c.CallCommand(ctx, "get_organization_admins", correlationId, data.NewAnyValueMapFromTuples(
		"org_id", orgId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *OrgRolesCommandableGrpcClientV1) GetOrganizationUserRoles(ctx context.Context, correlationId string, orgId string, paging *data.PagingParams) ([]*UserRolesV1, error) {
	res, err := c.CallCommand(ctx, "get_organization_user_roles", correlationId, data.NewAnyValueMapFromTuples(
		"org_id", orgId,
		"paging", paging,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]*UserRolesV1](res, correlationId)
}

func (c *OrgRolesCommandableGrpcClientV1) GrantOrgRole(ctx context.Context, correlationId string, orgId string, userId string, role string) ([]string, error) {
	res, err := c.CallCommand(ctx, "grant_org_role", correlationId, data.NewAnyValueMapFromTuples(
		"org_id", orgId,
		"user_id", userId,
		"user_role", role,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *OrgRolesCommandableGrpcClientV1) RevokeOrgRole(ctx context.Context, correlationId string, orgId string, userId string, role string) ([]string, error) {
	res, err := c.CallCommand(ctx, "revoke_org_role", correlationId, data.NewAnyValueMapFromTuples(
		"org_id", orgId,
		"user_id", userId,
		"user_role", role,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *OrgRolesCommandableGrpcClientV1) GrantDemoOrganizationUserRole(ctx context.Context, correlationId string, userId string, language string) (string, error) {
	res, err := c.CallCommand(ctx, "grant_demo_organization_user_role", correlationId, data.NewAnyValueMapFromTuples(
		"user_id", userId,
		"language", language,
	))

	if err != nil {
		return "", err
	}

	return clients.HandleHttpResponse[string](res, correlationId)
}
