package build

import (
	clients1 "github.com/pip-services-users2/client-orgroles-go/version1"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type OrgRolesClientFactory struct {
	*cbuild.Factory
}

func NewOrgRolesClientFactory() *OrgRolesClientFactory {
	c := &OrgRolesClientFactory{
		Factory: cbuild.NewFactory(),
	}

	cmdHttpClientDescriptor := cref.NewDescriptor("service-orgroles", "client", "commandable-http", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("service-orgroles", "client", "commandable-grpc", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-orgroles", "client", "grpc", "*", "1.0")

	c.RegisterType(cmdHttpClientDescriptor, clients1.NewOrgRolesCommandableHttpClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewOrgRolesCommandableGrpcClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewOrgRolesGrpcClientV1)

	return c
}
