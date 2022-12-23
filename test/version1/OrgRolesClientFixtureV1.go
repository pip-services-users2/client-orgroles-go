package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-users2/client-orgroles-go/version1"
	"github.com/stretchr/testify/assert"
)

type OrgRolesClientFixtureV1 struct {
	Client version1.IOrgRolesClientV1
}

func NewOrgRolesClientFixtureV1(client version1.IOrgRolesClientV1) *OrgRolesClientFixtureV1 {
	return &OrgRolesClientFixtureV1{
		Client: client,
	}
}

func (c *OrgRolesClientFixtureV1) TestGrantAndRevokeOrgRoles(t *testing.T) {

	// Grant organization role
	roles, err := c.Client.GrantOrgRole(context.Background(), "", "1", "123", "manager")
	assert.Nil(t, err)

	assert.NotNil(t, roles)
	assert.Len(t, roles, 1)
	assert.Equal(t, "1:manager", roles[0])

	// Get organization users
	_, err = c.Client.GetOrganizationUsers(context.Background(), "", "1")
	assert.Nil(t, err)

	// assert.Len(t, userIds, 4)

	// Revoke organization role
	roles, err = c.Client.RevokeOrgRole(context.Background(), "", "1", "123", "manager")
	assert.Nil(t, err)

	assert.Len(t, roles, 0)
}

func (c *OrgRolesClientFixtureV1) TestGrantDemoOrganizationUserRole(t *testing.T) {
	orgId, err := c.Client.GrantDemoOrganizationUserRole(context.Background(), "", "123", "en")

	assert.Nil(t, err)
	assert.NotEmpty(t, orgId)
}
