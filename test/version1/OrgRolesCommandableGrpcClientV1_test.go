package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-orgroles-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type orgRolesCommandableGrpcClientV1Test struct {
	client  *version1.OrgRolesCommandableGrpcClientV1
	fixture *OrgRolesClientFixtureV1
}

func newOrgRolesCommandableGrpcClientV1Test() *orgRolesCommandableGrpcClientV1Test {
	return &orgRolesCommandableGrpcClientV1Test{}
}

func (c *orgRolesCommandableGrpcClientV1Test) setup(t *testing.T) {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewOrgRolesCommandableGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewOrgRolesClientFixtureV1(c.client)
}

func (c *orgRolesCommandableGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableGrpcGrantAndRevokeOrgRoles(t *testing.T) {
	c := newOrgRolesCommandableGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestGrantAndRevokeOrgRoles(t)
}

func TestCommandableGrpcGrantDemoOrganizationUserRole(t *testing.T) {
	c := newOrgRolesCommandableGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestGrantDemoOrganizationUserRole(t)
}
