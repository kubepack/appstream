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

func (*MetadataServer) Get(ctx context.Context, req *api.MetadataGetRequest) (*api.MetadataGetResponse, error) {
	resp := &api.MetadataGetResponse{}
	resp.Status = serverutil.NewStatusOK()
	return resp, nil
}
