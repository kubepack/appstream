package app

import (
	api "github.com/appscode/appstream/pkg/apis/app/v1beta1"
	"github.com/appscode/appstream/pkg/apiserver/endpoints"
	"github.com/appscode/appstream/pkg/util"
	"golang.org/x/net/context"
)

func init() {
	endpoints.GRPCServerEndpoints.Register(api.RegisterMetadataServer, &MetadataServer{})
	endpoints.ProxyServerEndpoints.Register(api.RegisterMetadataHandlerFromEndpoint)
}

type MetadataServer struct{}

var _ api.MetadataServer = &MetadataServer{}

func (*MetadataServer) Git(ctx context.Context, req *api.GitRequest) (*api.GitResponse, error) {
	resp := &api.GitResponse{}
	resp.Status = serverutil.NewStatusOK()
	return resp, nil
}

func (*MetadataServer) Docker(ctx context.Context, req *api.DockerRequest) (*api.DockerResponse, error) {
	resp := &api.DockerResponse{}
	resp.Status = serverutil.NewStatusOK()
	return resp, nil
}
