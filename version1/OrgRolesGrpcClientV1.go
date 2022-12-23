package version1

import (
	"context"

	"github.com/pip-services-users2/client-orgroles-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type OrgRolesGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewOrgRolesGrpcClientV1() *OrgRolesGrpcClientV1 {
	return &OrgRolesGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("orgroles_v1.OrgRoles"),
	}
}

func (c *OrgRolesGrpcClientV1) GetOrganizationUsers(ctx context.Context, correlationId string, orgId string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "org_roles.get_organization_users")
	defer timing.EndTiming(ctx, err)

	req := &protos.OrgIdRequest{
		CorrelationId: correlationId,
		OrgId:         orgId,
	}

	reply := new(protos.UserIdsReply)
	err = c.CallWithContext(ctx, "get_organization_users", correlationId, req, reply)

	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	return reply.GetUserIds(), nil
}

func (c *OrgRolesGrpcClientV1) GetOrganizationAdmins(ctx context.Context, correlationId string, orgId string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "org_roles.get_organization_admins")
	defer timing.EndTiming(ctx, err)

	req := &protos.OrgIdRequest{
		CorrelationId: correlationId,
		OrgId:         orgId,
	}

	reply := new(protos.UserIdsReply)
	err = c.CallWithContext(ctx, "get_organization_admins", correlationId, req, reply)

	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	return reply.GetUserIds(), nil
}

func (c *OrgRolesGrpcClientV1) GetOrganizationUserRoles(ctx context.Context, correlationId string, orgId string, paging *data.PagingParams) (result []*UserRolesV1, err error) {
	timing := c.Instrument(ctx, correlationId, "org_roles.get_organization_user_roles")
	defer timing.EndTiming(ctx, err)

	req := &protos.RolesListRequest{
		CorrelationId: correlationId,
		OrgId:         orgId,
	}

	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.RolesListReply)
	err = c.CallWithContext(ctx, "get_organization_admins", correlationId, req, reply)

	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toUserRoles(reply.UserRoles)

	return result, nil
}

func (c *OrgRolesGrpcClientV1) GrantOrgRole(ctx context.Context, correlationId string, orgId string, userId string, role string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "org_roles.grant_org_role")
	defer timing.EndTiming(ctx, err)

	req := &protos.UserRoleRequest{
		CorrelationId: correlationId,
		OrgId:         orgId,
		UserId:        userId,
		UserRole:      role,
	}

	reply := new(protos.RolesReply)
	err = c.CallWithContext(ctx, "grant_org_role", correlationId, req, reply)

	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	return reply.Roles, nil
}

func (c *OrgRolesGrpcClientV1) RevokeOrgRole(ctx context.Context, correlationId string, orgId string, userId string, role string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "org_roles.revoke_org_role")
	defer timing.EndTiming(ctx, err)

	req := &protos.UserRoleRequest{
		CorrelationId: correlationId,
		OrgId:         orgId,
		UserId:        userId,
		UserRole:      role,
	}

	reply := new(protos.RolesReply)
	err = c.CallWithContext(ctx, "revoke_org_role", correlationId, req, reply)

	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	return reply.Roles, nil
}

func (c *OrgRolesGrpcClientV1) GrantDemoOrganizationUserRole(ctx context.Context, correlationId string, userId string, language string) (result string, err error) {
	timing := c.Instrument(ctx, correlationId, "org_roles.grant_demo_organization_user_role")
	defer timing.EndTiming(ctx, err)

	req := &protos.DemoOrganizationRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Language:      language,
	}

	reply := new(protos.OrgIdReply)
	err = c.CallWithContext(ctx, "grant_demo_organization_user_role", correlationId, req, reply)

	if err != nil {
		return "", err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return "", err
	}

	return reply.GetOrgId(), nil
}
