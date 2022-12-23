package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-orgroles-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type orgRolesCommandableHttpClientV1Test struct {
	client  *version1.OrgRolesCommandableHttpClientV1
	fixture *OrgRolesClientFixtureV1
}

func newOrgRolesCommandableHttpClientV1Test() *orgRolesCommandableHttpClientV1Test {
	return &orgRolesCommandableHttpClientV1Test{}
}

func (c *orgRolesCommandableHttpClientV1Test) setup(t *testing.T) {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewOrgRolesCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewOrgRolesClientFixtureV1(c.client)
}

func (c *orgRolesCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpGrantAndRevokeOrgRoles(t *testing.T) {
	c := newOrgRolesCommandableHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestGrantAndRevokeOrgRoles(t)
}

func TestCommandableHttpGrantDemoOrganizationUserRole(t *testing.T) {
	c := newOrgRolesCommandableHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestGrantDemoOrganizationUserRole(t)
}
