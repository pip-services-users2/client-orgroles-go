package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-orgroles-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type orgRolesGrpcClientV1Test struct {
	client  *version1.OrgRolesGrpcClientV1
	fixture *OrgRolesClientFixtureV1
}

func newOrgRolesGrpcClientV1Test() *orgRolesGrpcClientV1Test {
	return &orgRolesGrpcClientV1Test{}
}

func (c *orgRolesGrpcClientV1Test) setup(t *testing.T) {
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

	c.client = version1.NewOrgRolesGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewOrgRolesClientFixtureV1(c.client)
}

func (c *orgRolesGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGrpcGrantAndRevokeOrgRoles(t *testing.T) {
	c := newOrgRolesGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestGrantAndRevokeOrgRoles(t)
}

func TestGrpcGrantDemoOrganizationUserRole(t *testing.T) {
	c := newOrgRolesGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestGrantDemoOrganizationUserRole(t)
}
